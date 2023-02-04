import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-statistics',
  templateUrl: './statistics.component.html',
  styleUrls: ['./statistics.component.css']
})
export class StatisticsComponent implements OnInit {

  allTime = 0;
  today = 0;
  getAllTimeAddedUrl = "http://localhost:8080/api/receipts/allcount";
  getTodayAddedUrl = "http://localhost:8080/api/receipts/allcounttoday";

  constructor(private http: HttpClient){

  }

  ngOnInit(): void {
    this.http.get(this.getAllTimeAddedUrl, {withCredentials: false}).subscribe(
      (res: any) => {
        this.allTime = res;
      },
      err => {
        console.error(err);        
      }
    );
    this.http.get(this.getTodayAddedUrl, {withCredentials: false}).subscribe(
      (res: any) => {
        this.today = res;
      },
      err => {
        console.error(err);        
      }
    );
  }

}
