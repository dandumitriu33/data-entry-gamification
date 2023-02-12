import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FileDownloadService {

  private avatarDownloadURL = "http://localhost:8080/api/user/avatar";

  constructor(private httpClient: HttpClient) { }

  downloadAvatar(): Observable<Blob> {
    return this.httpClient.get(this.avatarDownloadURL, {
      withCredentials: true,
      responseType: 'blob'
    })
  }
}
