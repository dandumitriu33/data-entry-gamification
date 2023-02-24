import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { Emitters } from 'src/app/emitters/emitters';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  form: FormGroup;
  loginUrl = "http://localhost:8080/api/login";

  constructor(private formBuilder: FormBuilder,
    private http: HttpClient,
    private router: Router) {
  }

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      email: '',
      password: ''
    })
  }

  submit(): void {
    this.http.post(this.loginUrl, this.form.getRawValue(), {
      withCredentials: true
    }).subscribe(()=> this.router.navigate(['/']));
  }
}
