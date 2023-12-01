import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

export interface appresponse{
  AP : number
  DID: number
  Date: string
  DoctorFirstName: string
  DoctorLastName: string
  Hour: string
  PID: number
  SID: number
 
}

export interface Slots {
  slot_id: number
  date: string
  hour: string
  empty: boolean
}

@Injectable({
  providedIn: 'root'
})
export class PatientService {

  constructor(private httpClient: HttpClient) {}

  saveUser(userData: object){
    return this.httpClient.post('http://localhost:8080/signup' , userData);
  }
  signinfunc(userData: object){
    return this.httpClient.post('http://localhost:8080/signin' , userData);
  }
  addslot(drdata: object){
    return this.httpClient.post('http://localhost:8080/addslot' , drdata);
  }
  getappserv(){
    return this.httpClient.get('http://localhost:8080/getuserappointments');
  }
  onReserve(payload: object){
    return this.httpClient.post('http://localhost:8080/makeappointment' , payload);
  }
  getslotserv(){
    return this.httpClient.get('http://localhost:8080/getdoctorslotsbyid');
  }
  deleteapp(userData: any) {
    return this.httpClient.post('http://localhost:8080/cancelappointment', userData);
  }
  updateapp(userData: any) {
  return this.httpClient.post('http://localhost:8080/cancelappointment', userData);
  }


}
