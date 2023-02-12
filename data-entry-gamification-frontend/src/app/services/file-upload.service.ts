import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpEvent } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FileUploadService {

  private avatarUploadURL = "http://localhost:8080/api/user/avatar";

  constructor(private http: HttpClient) {}

  upload(file: File): Observable<HttpEvent<any>> {
    const formData: FormData = new FormData();

    formData.append('file', file);

    const req = new HttpRequest('PUT', this.avatarUploadURL, formData, {
      reportProgress: true,
      responseType: 'json',
    });

    return this.http.request(req);
  }

  getFiles(): Observable<any> {
    // return this.http.get(this.avatarUploadURL);
    return this.http.get("http://localhost:8080/api/user/info");
  }
}
