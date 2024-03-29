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
  userLoggedIn = false;
  logoutUrl = "http://localhost:8080/api/logout";
  getUserUrl = "http://localhost:8080/api/user";
  getUserRolesUrl = "http://localhost:8080/api/user/roles";
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
    this.http.get(this.getUserRolesUrl, {withCredentials: true}).subscribe(
      (res: any) => {
        console.log(res)
        this.roles = res;
        console.log(this.roles)
      },
      err => {
        console.error(err);
        console.log("navbar emitter - auth false");
      }
    );
    
    Emitters.authEmitter.subscribe(
      (auth: boolean) => {
        this.authenticated = auth;
      }
    );
    Emitters.loginEmitter.subscribe(
      (auth: boolean) => {
        this.userLoggedIn = auth;
        this.http.get(this.getUserRolesUrl, {withCredentials: true}).subscribe(
          (res: any) => {
            console.log(res)
            this.roles = res;
            console.log(this.roles)
          },
          err => {
            console.error(err);
            console.log("navbar emitter - auth false");
          }
        );
      }
    );
    Emitters.logoutEmitter.subscribe(
      (auth: boolean) => {
        if (auth == true) {
          this.userLoggedIn = false;
        }
      }
    );
  }

  
  logout(): void {
    this.roles = [];
    this.http.get(this.logoutUrl, {withCredentials: true})
    .subscribe(() => {
      this.authenticated = false;
      Emitters.authEmitter.emit(false);
      Emitters.logoutEmitter.emit(true);
    });
  }
  
  rolesContainsQA(): boolean {
    if (this.roles == null) return false
    if (this.roles.indexOf("qa") != -1 && this.authenticated) {
      return true
    }
    return false
  }
}
