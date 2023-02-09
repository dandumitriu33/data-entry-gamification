import { Component, OnInit } from '@angular/core';
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

  constructor(private http: HttpClient) {

  }

  ngOnInit(): void {
    this.refreshPoints(); 
    Emitters.inputEmitter.subscribe(
      (input: boolean) => {
        console.log("refreshing points input", input)
        this.refreshPoints()
      }
    );
    Emitters.authEmitter.subscribe(
      (auth: boolean) => {
        console.log("refreshing points")
        this.refreshPoints()
      }
    );
  }  

  refreshPoints(): void {
    console.log("refreshing points func")
    this.http.get(this.getUserInfoUrl, {withCredentials: true}).subscribe(
      (res: any) => {
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
