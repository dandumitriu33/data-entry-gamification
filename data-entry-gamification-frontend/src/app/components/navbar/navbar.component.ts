import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Emitters } from 'src/app/emitters/emitters';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  authenticated = false;
  logoutUrl = "http://localhost:8080/api/logout";
  getUserUrl = "http://localhost:8080/api/user";
  roles: string[] = [];

  constructor(private http: HttpClient) {

  }

  ngOnInit(): void {
    this.http.get(this.getUserUrl, {withCredentials: true}).subscribe(
      (res: any) => {
        console.log("navbar emitter - auth true");
        Emitters.authEmitter.emit(true);
      },
      err => {
        console.error(err);
        console.log("navbar emitter - auth false");
        Emitters.authEmitter.emit(false);
      }
    );
    Emitters.authEmitter.subscribe(
      (auth: boolean) => {
        this.authenticated = auth;
      }
    );
  }

  
  logout(): void {
    this.http.get(this.logoutUrl, {withCredentials: true})
    .subscribe(() => {
      this.authenticated = false;
      Emitters.authEmitter.emit(false);
    });
  }
  
  rolesContainsQA(): boolean {
    if (this.roles.indexOf("qa") != -1) {
      return true
    }
    return false
  }
}
