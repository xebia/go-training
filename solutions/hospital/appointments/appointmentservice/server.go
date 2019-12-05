package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentapi"
	"github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"
	"github.com/MarcGrol/go-training/solutions/hospital/patients/patientinfoapi"
)

type server struct {
	listener net.Listener
	pb.UnimplementedAppointmentExternalServer
	pb.UnimplementedAppointmentInternalServer

	appointmentStore    AppointmentStore
	patientInfoService  patientinfoapi.PatientInfoServer
	notificationService notificationapi.NotificationServer
}

func newServer(store AppointmentStore) *server {
	return &server{
		appointmentStore: store,
	}
}

func (s *server) GRPCListenBlocking(port string) error {
	var err error
	s.listener, err = net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen at port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAppointmentExternalServer(grpcServer, newServer(s.appointmentStore))
	pb.RegisterAppointmentInternalServer(grpcServer, newServer(s.appointmentStore))

	log.Println("GRPPC server starts listening...")
	err = grpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (s *server) GetAppointmentsOnUser(c context.Context, in *pb.GetAppointmentsOnUserRequest) (*pb.GetAppointmentsReply, error) {
	// Validate input
	// TODO

	// Perform lookup
	internalAppointments, err := s.appointmentStore.GetAppointmentsOnUserUid(in.UserUid)
	if err != nil {
		return &pb.GetAppointmentsReply{
			Error: &pb.Error{
				Code:    500,
				Message: "Technical error fetching appointments on user",
				Details: err.Error(),
			},
		}, nil
	}

	return returnAppointmentList(internalAppointments), nil
}

func (s *server) GetAppointmentsOnStatus(c context.Context, in *pb.GetAppointmentsOnStatusRequest) (*pb.GetAppointmentsReply, error) {
	// Validate input
	// TODO

	// Perform lookup
	internalAppointments, err := s.appointmentStore.GetAppointmentsOnStatus(AppointmentStatus(in.Status))
	if err != nil {
		return &pb.GetAppointmentsReply{
			Error: &pb.Error{
				Code:    500,
				Message: "Technical error fetching appointments on user",
				Details: err.Error(),
			},
		}, nil
	}

	return returnAppointmentList(internalAppointments), nil
}

func (s *server) RequestAppointment(c context.Context, in *pb.RequestAppointmentRequest) (*pb.AppointmentReply, error) {
	// Validate input
	// TODO

	// Adjust datastore
	appointmentCreated, err := s.appointmentStore.PutAppointment(convertIntoInternal(*in.Appointment))
	if err != nil {
		return &pb.AppointmentReply{
			Error: &pb.Error{
				Code:    500,
				Message: "Technical error creating appointment",
				Details: err.Error(),
			},
		}, nil
	}
	return returnSingleAppointment(appointmentCreated), nil
}

func (s *server) ModifyAppointmentStatus(c context.Context, in *pb.ModifyAppointmentStatusRequest) (*pb.AppointmentReply, error) {
	// Validate input
	// TODO

	// Perform lookup
	internalAppointment, found, err := s.appointmentStore.GetAppointmentOnUid(in.AppointmentUid)
	if err != nil {
		return &pb.AppointmentReply{
			Error: &pb.Error{
				Code:    500,
				Message: "Error persisting new appointment",
				Details: err.Error(),
			},
		}, nil
	}
	if !found {
		return &pb.AppointmentReply{
			Error: &pb.Error{
				Code:    404,
				Message: "Appointment with uid not found",
			},
		}, nil
	}

	// Fetch patient details
	resp, err := s.patientInfoService.GetPatientOnUid(c, &patientinfoapi.GetPatientOnUidRequest{PatientUid: internalAppointment.UserUID})
	if err != nil {
		return &pb.AppointmentReply{
			Error: &pb.Error{
				Code:    500,
				Message: "Technical error finding patient on uid",
				Details: err.Error(),
			},
		}, nil
	}
	if resp.Error != nil {
		return &pb.AppointmentReply{
			Error: &pb.Error{
				Code:    resp.Error.Code,
				Message: fmt.Sprintf("Error getting patient on uid:%s", resp.Error.Message),
				Details: resp.Error.Details,
			},
		}, nil
	}

	// Send out sms
	_, _ = s.notificationService.SendSms(c, &notificationapi.SendSmsRequest{
		Sms: &notificationapi.SmsMessage{
			RecipientPhoneNumber: resp.Patient.PhoneNumber,
			Body:                 "Appointment confirmed", // TODO use template
		},
	})
	// TODO error checking

	// Send out email
	_, _ = s.notificationService.SendEmail(c, &notificationapi.SendEmailRequest{
		Email: &notificationapi.EmailMessage{
			RecipientEmailAddress: resp.Patient.PhoneNumber,
			Subject:               "Appointment confirmed",              // TODO use template
			Body:                  "Appointment confirmed with details", // TODO use template
		},
	})
	// TODO error checking

	// Adjust datastore
	internalAppointment.Status = AppointmentStatusConfirmed
	appointmentAdjusted, err := s.appointmentStore.PutAppointment(internalAppointment)
	return &pb.AppointmentReply{
		Error: &pb.Error{
			Code:    500,
			Message: "Error persisting modified appointment",
			Details: resp.Error.Details,
		},
	}, nil

	return returnSingleAppointment(appointmentAdjusted), nil
}

func returnAppointmentList(internalAppointments []Appointment) *pb.GetAppointmentsReply {
	externalAppointments := []*pb.Appointment{}
	for _, a := range internalAppointments {
		externalAppointments = append(externalAppointments, convertIntoExternal(a))
	}
	return &pb.GetAppointmentsReply{
		Appointments: externalAppointments,
	}
}

func returnSingleAppointment(internalAppointmnent Appointment) *pb.AppointmentReply {
	return &pb.AppointmentReply{
		Appointment: convertIntoExternal(internalAppointmnent),
	}
}

func convertIntoExternal(a Appointment) *pb.Appointment {
	return &pb.Appointment{
		AppointmentUid: a.AppointmentUID,
		UserUid:        a.UserUID,
		DateTime:       a.DateTime,
		Location:       a.Location,
		Topic:          a.Topic,
		Status:         pb.AppointmentStatus(a.Status),
	}
}

func convertIntoInternal(a pb.Appointment) Appointment {
	return Appointment{
		AppointmentUID: a.AppointmentUid,
		UserUID:        a.UserUid,
		DateTime:       a.DateTime,
		Location:       a.Location,
		Topic:          a.Topic,
		Status:         AppointmentStatus(a.Status),
	}
}