import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { AvatarComponent } from './components/avatar/avatar.component';
import { PaperSimComponent } from './components/paper-sim/paper-sim.component';
import { HomeComponent } from './components/home/home.component';
import { InputComponent } from './components/input/input.component';
import { ProfileComponent } from './components/profile/profile.component';
import { ReceiptComponent } from './components/receipt/receipt.component';
import { HttpClientModule } from '@angular/common/http';
import { ReceiptFormComponent } from './components/receipt-form/receipt-form.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RegisterComponent } from './components/register/register.component';
import { LoginComponent } from './components/login/login.component';
import { StatisticsComponent } from './components/statistics/statistics.component';

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    AvatarComponent,
    PaperSimComponent,
    HomeComponent,
    InputComponent,
    ProfileComponent,
    ReceiptComponent,
    ReceiptFormComponent,
    RegisterComponent,
    LoginComponent,
    StatisticsComponent,
  ],
  imports: [
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
