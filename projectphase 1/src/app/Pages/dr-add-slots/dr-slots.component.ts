import { Component } from '@angular/core';
import { PatientService } from '../Services/patient.service';

@Component({
  selector: 'app-dr-slots',
  templateUrl: './dr-slots.component.html',
  styleUrls: ['./dr-slots.component.css']
})
export class DrSlotsComponent {
  constructor(private PatientService: PatientService){}
  selectedDate!: string;
  selectedTime!: string;
 

  onDateChange(event: any) {
    this.selectedDate = event.target.value;
  }

  onTimeChange(event: any) {
    this.selectedTime = event.target.value;
  }



  addslot() {
   
    const drdata = {
      date: this.selectedDate,
      hour: this.selectedTime
    };
    console.log(drdata);
    console.log(drdata)
    // Make an HTTP request to save the user data
    this.PatientService.addslot(drdata).subscribe({
      next: (res: any) =>{
        console.log(res,'response')

      },
      error: (err:any) =>{
        console.log(err,'errors')

      }
    })
    
  }
}
