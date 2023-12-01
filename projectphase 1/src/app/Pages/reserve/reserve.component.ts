import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { PatientService } from '../Services/patient.service';

interface User {
  user_id: number;
  first_name: string;
  last_name: string;
}

interface Slot {
  slot_id: number;
  date: string;
  hour: string;
  empty: boolean;
}

@Component({
  selector: 'app-reserve',
  templateUrl: './reserve.component.html',
  styleUrls: ['./reserve.component.css']
})
export class ReserveComponent implements OnInit {
  selectedDoctor: number | null;
  selectedAppointment: Slot = { slot_id: 0, date: '', hour: '', empty: false };
  doctors: User[] = [];
  appointmentOptions: Slot[] = [];

  constructor(private http: HttpClient , private PatientService: PatientService) {
    this.selectedDoctor = null;
    this.selectedAppointment;
  }

  ngOnInit() {
    // Fetch doctors data from the API endpoint
    this.http.get<User[]>('http://localhost:8080/getdoctors').subscribe((data) => {
      this.doctors = data;
    });
  }

  onSelectedDoctorChange() {
    // Fetch appointment options from the API endpoint based on the selected doctor
    if (this.selectedDoctor) {
      this.http.get<Slot[]>('http://localhost:8080/getdoctorslots?doctor_id=' + this.selectedDoctor).subscribe((data) => {
        this.appointmentOptions = data;
      });
    } else {
      // Clear the appointment options if no doctor is selected
      this.appointmentOptions = [];
    }
  }

  onSelectedappChange() {
    // Handle changes when the user selects an appointment
    // You can access the selected slot_id using this.selectedAppointment.slot_id
    if (this.selectedAppointment) {
      console.log(this.selectedAppointment.slot_id);
    }
  }

  onReserve() {
    // Call the backend API to make the reservation using this.selectedAppointment.slot_id
    const payload = {
      slot_id: Number(this.selectedAppointment.slot_id),
      
    };
      
      console.log(payload)
      this.PatientService.onReserve(payload).subscribe({
        next: (res: any) =>{
          console.log(res,'response')
          
        },
        error: (err:any) =>{
          console.log(err,'errors')
  
        }
      })
    
  }
}
