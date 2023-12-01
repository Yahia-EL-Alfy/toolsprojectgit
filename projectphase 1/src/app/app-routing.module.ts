import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './Pages/home-page/home-page.component';
import { SignUpComponent } from './Pages/sign-up/sign-up.component';
import { SignInComponent } from './Pages/sign-in/sign-in.component';
import { DrSignInComponent } from './Pages/dr-sign-in/dr-sign-in.component';
import { DrSlotsComponent } from './Pages/dr-add-slots/dr-slots.component';
import { ViewDrSlotsComponent } from './Pages/view-dr-slots/view-dr-slots.component';
import { ViewAppComponent } from './Pages/view-app/view-app.component';
import { ReserveComponent } from './Pages/reserve/reserve.component';

const routes: Routes = [
  { path: '', component: HomePageComponent , title : 'Home'},
  { path: 'sign-in', component: SignInComponent , title : 'Sign-in'},
  { path: 'sign-up', component: SignUpComponent , title : 'Sign-up'},
  { path: 'dr-sign-in', component: DrSignInComponent , title : 'Dr Sign-in'},
  { path: 'doctor/dr-add-slots', component: DrSlotsComponent , title : 'Dr Add Slots'},
  { path: 'doctor/dr-view-slots', component: ViewDrSlotsComponent , title : 'Dr View Slots'},
  { path: 'patient/view-appointments', component: ViewAppComponent , title : 'View Appointment '},
  { path: 'patient/reserve', component: ReserveComponent , title : 'Reserve'},

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
