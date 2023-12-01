import { Component } from '@angular/core';
import { PatientService, appresponse } from '../Services/patient.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-view-app',
  templateUrl: './view-app.component.html',
  styleUrls: ['./view-app.component.css']
})
export class ViewAppComponent {
  constructor(private PatientService: PatientService, private router: Router) {}

  app!:appresponse[];

  ngOnInit() {

    this.getapplist();
  }
  getapplist(){
    this.PatientService.getappserv().subscribe((res:any)=>{
      console.log(res)
      this.app=res
    });
  }
  deleteapp(event: any, AP: number) {
    const userData = {
        appointment_id: AP
    };

    console.log(userData);

    // Make an HTTP request to cancel the appointment
    this.PatientService.deleteapp(userData).subscribe({
        next: (res: any) => {
            console.log(res, 'response');
            // Handle success
        },
        error: (err: any) => {
            console.log(err, 'errors');
            // Handle errors
        }
    });
}
updateapp(event: any, AP: number) {
  const userData = {
    appointment_id: AP
  };

  console.log(userData);
this.router.navigate(['/patient/view-appointments']);
  // Make an HTTP request to update the appointment
  this.PatientService.updateapp(userData).subscribe({
    next: (res: any) => {
      console.log(res, 'response');
      // Handle success

      // Navigate to the /patient/view-appointments route
      this.router.navigate(['/patient/view-appointments']);
    },
    
    error: (err: any) => {
      this.router.navigate(['/patient/view-appointments']);
      console.log(err, 'errors');
      // Handle errors
    }
  });
}
}
