import { Component, OnInit } from '@angular/core';
import { Emitters } from 'src/app/emitters/emitters';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-avatar',
  templateUrl: './avatar.component.html',
  styleUrls: ['./avatar.component.css']
})
export class AvatarComponent implements OnInit {
  userPoints = "0"; // points over 1k will be dislpayed as 1.2 K
  userLevel = 0;
  getPointsUrl = "";

  constructor(private http: HttpClient) {

  }

  ngOnInit(): void {
    this.refreshPoints(); 
    Emitters.inputEmitter.subscribe(
      (auth: boolean) => {
        // refresh points
        console.log("refreshing points")
      }
    );
  }  

  refreshPoints(): void {
    this.http.get(this.getPointsUrl, {withCredentials: true}).subscribe(
      (res: any) => {
        console.log("got points for user");
        // TODO: adjust after endpoint creation
        this.userPoints = this.formatNumberToString(Math.floor( Math.random() * 7000 ))
      },
      err => {
        console.error(err);    
        // TODO: remove after endpoint creation 
        this.userPoints = this.formatNumberToString(Math.floor( Math.random() * 7000 ))
      }
    );
  };

  formatNumberToString(num: number): string {
    let result = 0;
    let printResult = "0";
    if (num >= 1000000) {
      printResult = (num / 1000000).toFixed(2) + "M";
    } else if (num >= 1000) {
      printResult = (num / 1000).toFixed(2) + "K";
    } else {
      printResult = result.toString();
    }
    return printResult;
  }
}
