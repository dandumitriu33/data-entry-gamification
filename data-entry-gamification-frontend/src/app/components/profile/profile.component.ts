import { Component, OnInit } from '@angular/core';
import { UserInfo } from 'src/app/entities/user_info';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  userPoints = 0;
  userLevel = 0;
  getUserInfoUrl = "http://localhost:8080/api/user/info";

  constructor(private http: HttpClient){

  }

  ngOnInit(): void {
    this.http.get(this.getUserInfoUrl, {withCredentials: true}).subscribe(
      (res: any) => {
        this.userPoints = res.points;
        this.userLevel = res.level;
      },
      err => {
        console.error(err);        
      }
    );
  }

}
