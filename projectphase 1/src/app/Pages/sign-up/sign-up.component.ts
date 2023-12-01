import { Component } from '@angular/core';
import { PatientService } from '../Services/patient.service';

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent {
  constructor(private PatientService: PatientService){}
  first_name!: string;
  last_name!: string;
  email!: string;
  password!: string;
  is_doctor!: boolean;



  saveUser() {
    // Create an object with user data
    const userData = {
      first_name: this.first_name,
      last_name: this.last_name,
      email: this.email,
      password: this.password,
      is_doctor: this.is_doctor
    };

    console.log(userData)
    // Make an HTTP request to save the user data
    this.PatientService.saveUser(userData).subscribe({
      next: (res: any) =>{
        console.log(res,'response')

      },
      error: (err:any) =>{
        console.log(err,'errors')

      }
    })
    
  }
}
