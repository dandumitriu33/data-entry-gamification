import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { Emitters } from 'src/app/emitters/emitters';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-avatar',
  templateUrl: './avatar.component.html',
  styleUrls: ['./avatar.component.css'],
})
export class AvatarComponent implements OnInit {
  userPoints = "0"; // points over 1k will be dislpayed as 1.2 K
  userLevel = 0;
  getUserInfoUrl = "http://localhost:8080/api/user/info";
  // avatarURI = "/assets/img/001-Default-Avatar.jpg"
  avatarURI = "http://localhost:8080/api/user/avatar?t=" + new Date().getTime();

  constructor(private http: HttpClient,
              private cdr: ChangeDetectorRef) {

  }

  ngOnInit(): void {
    this.refreshPoints(); 
    Emitters.inputEmitter.subscribe(
      (input: boolean) => {
        console.log("avatar refreshing points input")
        setTimeout(() => {
          console.log("input point refresh 5s")
          this.refreshPoints()
        }, 5000); 
      }
    );
    Emitters.authEmitter.subscribe(
      (auth: boolean) => {
        console.log("avatar refreshing points auth")
        this.refreshPoints()
        this.avatarURI = "/assets/img/loading.gif"
        console.log("detect changes 1", this.avatarURI);
        this.cdr.detectChanges();
        setTimeout(() => {
          this.avatarURI = "http://localhost:8080/api/user/avatar?t=" + new Date().getTime();
          console.log("detect changes 2", this.avatarURI);
          this.cdr.detectChanges();
        }, 1000);
        
      }
    );
    Emitters.uploadAvatarEmitter.subscribe(
      (auth: boolean) => {
        console.log("avatar refreshing image auth")
        this.avatarURI = "/assets/img/loading.gif";
        setTimeout(() => this.avatarURI = "http://localhost:8080/api/user/avatar", 1000);
      }
    );
  }  

  refreshPoints(): void {
    this.http.get(this.getUserInfoUrl, {withCredentials: true}).subscribe(
      (res: any) => {
        console.log(res)
        this.userPoints = this.formatNumberToString(res.points);
        this.userLevel = res.level;
      },
      err => {
        console.error(err);        
      }
    );    
  };

  formatNumberToString(num: number): string {
    let printResult = "0";
    if (num >= 1000000) {
      printResult = (num / 1000000).toFixed(2) + "M";
    } else if (num >= 1000) {
      printResult = (num / 1000).toFixed(2) + "K";
    } else {
      printResult = num.toString();
    }
    return printResult;
  }
}
