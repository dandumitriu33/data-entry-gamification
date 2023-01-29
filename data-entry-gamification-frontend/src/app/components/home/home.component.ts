import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Emitters } from 'src/app/emitters/emitters';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  message = '';
  getUserUrl = "http://localhost:8080/api/user";

  constructor(private http: HttpClient){

  }

  ngOnInit(): void {
    this.http.get(this.getUserUrl, {withCredentials: true}).subscribe(
      (res: any) => {
        this.message = `Hello ${res.first_name}`;
        Emitters.authEmitter.emit(true);
      },
      err => {
        console.error(err);
        this.message = "You are not logged in."
        Emitters.authEmitter.emit(false);
      }
    );
  }
}
