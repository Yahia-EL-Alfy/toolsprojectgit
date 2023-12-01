import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomePageComponent } from './Pages/home-page/home-page.component';
import { SignInComponent } from './Pages/sign-in/sign-in.component';
import { SignUpComponent } from './Pages/sign-up/sign-up.component';
import { DrSignInComponent } from './Pages/dr-sign-in/dr-sign-in.component';
import { DrSlotsComponent } from './Pages/dr-add-slots/dr-slots.component';
import { FormsModule } from '@angular/forms';
import { ViewDrSlotsComponent } from './Pages/view-dr-slots/view-dr-slots.component';
import { ViewAppComponent } from './Pages/view-app/view-app.component';
import { ReserveComponent } from './Pages/reserve/reserve.component';


@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    SignInComponent,
    SignUpComponent,
    DrSignInComponent,
    DrSlotsComponent,
    ViewDrSlotsComponent,
    ViewAppComponent,
    ReserveComponent,
    
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
