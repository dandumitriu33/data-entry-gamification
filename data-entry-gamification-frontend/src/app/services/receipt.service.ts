import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { InterfaceReceipt } from '../interfaces/interface-receipt';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
import { Receipt, ReceiptDTO } from '../entities/receipt';

@Injectable({
  providedIn: 'root'
})
export class ReceiptService {
  private receiptsUrl = 'http://localhost:8080/api/receipts';  // URL to web api
  private updateVerifiedReceiptURL = "http://localhost:8080/api/receipts/verified";
  receiptDTO: ReceiptDTO = {
    id: 0, 
    model_year: 0, 
    make: "", 
    vin: "", 
    first_name: "", 
    last_name: "", 
    state: "", 
    date_added: "", 
    qa_score: 0, 
    qa_date: ""
  };

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
    withCredentials: true
  };

  constructor(private http: HttpClient) { }

  /** GET receipt by id. Will 404 if id not found */
  getReceipt(id: number): Observable<InterfaceReceipt> {
    const url = `${this.receiptsUrl}/${id}`;
    return this.http.get<InterfaceReceipt>(url).pipe(
      tap(_ => console.info(`fetched receipt id=${id}`)),
      catchError(this.handleError<InterfaceReceipt>(`getReceipt id=${id}`))
    );
  }

  /** POST: add a new receipt*/
  addReceipt(receipt: Receipt): Observable<ReceiptDTO> {
    // map receipt to ReceiptDTO
    this.receiptDTO.id = receipt.id;
    this.receiptDTO.model_year = receipt.model_year, 
    this.receiptDTO.make = receipt.make, 
    this.receiptDTO.vin = receipt.vin, 
    this.receiptDTO.first_name = receipt.first_name, 
    this.receiptDTO.last_name = receipt.last_name, 
    this.receiptDTO.state = receipt.state, 
    this.receiptDTO.date_added = receipt.date_added, 
    this.receiptDTO.qa_score = receipt.qa_score, 
    this.receiptDTO.qa_date = receipt.qa_date
    
    return this.http.post<ReceiptDTO>(this.receiptsUrl, this.receiptDTO, this.httpOptions).pipe(
      tap((newReceipt: ReceiptDTO) => console.info(`added receipt w/ id=${newReceipt.id}`)),
      catchError(this.handleError<ReceiptDTO>('addReceipt'))
    );
  }

  /** PUT: update a rceipt with QA Score and Date */
  updateVerifiedReceipt(receiptDTO: ReceiptDTO): Observable<ReceiptDTO> {
    return this.http.put<ReceiptDTO>(this.updateVerifiedReceiptURL, receiptDTO, this.httpOptions).pipe(
      tap((newReceipt: ReceiptDTO) => console.info(`updated receipt w/ id=${newReceipt.id}`)),
      catchError(this.handleError<ReceiptDTO>('updateVerifiedReceipt'))
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
