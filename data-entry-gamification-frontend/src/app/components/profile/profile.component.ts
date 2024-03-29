import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Emitters } from 'src/app/emitters/emitters';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  userPoints = 0;
  userLevel = 0;
  getUserInfoUrl = "http://localhost:8080/api/user/info";
  downloadUserAvatarURL = "http://localhost:8080/api/user/avatar";

  constructor(private http: HttpClient){

  }

  ngOnInit(): void {
    this.refreshData();
    Emitters.uploadAvatarEmitter.subscribe(
      () => {
        this.downloadUserAvatarURL = "";
        setTimeout(() => this.downloadUserAvatarURL = "http://localhost:8080/api/user/avatar", 1000);
      }
    );
    Emitters.authEmitter.subscribe(
      (auth: boolean) => {
        this.downloadUserAvatarURL = "";
        setTimeout(() => this.downloadUserAvatarURL = "http://localhost:8080/api/user/avatar", 1000);
      }
    );
  }

  refreshData() {
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
