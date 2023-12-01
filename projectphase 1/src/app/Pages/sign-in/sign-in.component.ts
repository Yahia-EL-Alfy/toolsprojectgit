import { Component } from '@angular/core';
import { PatientService } from '../Services/patient.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.css']
})
export class SignInComponent {
  constructor(private PatientService: PatientService , private router: Router){}
  email!: string;
  password!: string;


  signinfunc() {
    const userData = {
      email: this.email,
      password: this.password,
    };

    this.PatientService.signinfunc(userData).subscribe({
      next: (res: any) => {
        // Assuming your API response contains a field named 'is_doctor'
        const isDoctor = res.is_doctor;

        if (isDoctor) {
          // User is a doctor, navigate to doctor's page
          this.router.navigate(['/doctor/dr-view-slots']);
        } else {
          // User is not a doctor, navigate to patient's page
          this.router.navigate(['/patient/view-appointments']);
        }
      },
      error: (err: any) => {
        console.log(err, 'errors');
        // Handle errors if needed
      },
    });
  }
}
