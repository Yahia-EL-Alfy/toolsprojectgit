import { Component } from '@angular/core';
import { PatientService, Slots } from '../Services/patient.service';

@Component({
  selector: 'app-view-dr-slots',
  templateUrl: './view-dr-slots.component.html',
  styleUrls: ['./view-dr-slots.component.css']
})
export class ViewDrSlotsComponent {
  constructor(private PatientService: PatientService){}

  slots!: Slots[];

  ngOnInit() {

    this.getslotList();
  }
  getslotList(){
    this.PatientService.getslotserv().subscribe((res:any)=>{
      console.log(res)
      this.slots=res
    });
  }
}
