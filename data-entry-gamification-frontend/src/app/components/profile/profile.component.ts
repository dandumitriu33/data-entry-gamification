import { Component, OnInit } from '@angular/core';
import { UserInfo } from 'src/app/entities/user_info';
import { HttpClient } from '@angular/common/http';
import { FileDownloadService } from 'src/app/services/file-download.service';

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
  }

  refreshData() {
    this.http.get(this.getUserInfoUrl, {withCredentials: true}).subscribe(
      (res: any) => {
        this.userPoints = res.points;
        this.userLevel = res.level;
        console.log("user avatar location:", res.img_uri);
      },
      err => {
        console.error(err);        
      }
    );
  }

}
