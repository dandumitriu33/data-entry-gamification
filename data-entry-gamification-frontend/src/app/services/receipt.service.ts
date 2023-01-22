import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Receipt } from '../interfaces/receipt';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ReceiptService {
  private receiptsUrl = 'http://localhost:8080/receipts';  // URL to web api
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(private http: HttpClient) { }

  /** GET receipt by id. Will 404 if id not found */
  getReceipt(id: number): Observable<Receipt> {
    const url = `${this.receiptsUrl}/${id}`;
    return this.http.get<Receipt>(url).pipe(
      tap(_ => console.info(`fetched receipt id=${id}`)),
      catchError(this.handleError<Receipt>(`getReceipt id=${id}`))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   *
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {     
      // TODO: better job of transforming error for user consumption
      console.error(error); // log to console 

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
}
