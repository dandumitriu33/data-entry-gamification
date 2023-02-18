import { Component, OnInit, ElementRef, ViewChild } from '@angular/core';
import { Receipt, ReceiptDTO } from 'src/app/entities/receipt';
import { ReceiptService } from 'src/app/services/receipt.service';
import { Emitters } from 'src/app/emitters/emitters';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-qa',
  templateUrl: './qa.component.html',
  styleUrls: ['./qa.component.css']
})
export class QaComponent implements OnInit {
  @ViewChild("modelYear", {  }) modelYear: ElementRef;

  getLatestUnverifiedReceiptURL = "http://localhost:8080/api/receipts/unverified";

  receipt: Receipt = {
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

  receiptDTO: ReceiptDTO = {
                      id: 0, 
                      model_year: 0, 
                      make: "", 
                      vin: "", 
                      first_name: "", 
                      last_name: "", 
                      state: "", 
                      date_added: {String: "", Valid: true}, 
                      qa_score: {Int64: 0, Valid: false}, 
                      qa_date: {String: "", Valid: false}
  } 

  constructor(
    private receiptService: ReceiptService,
    private http: HttpClient
  ) {}

  ngOnInit(): void {
    this.refreshData();
    Emitters.inputEmitter.subscribe(
      () => {
        this.modelYear.nativeElement.focus()
        console.log("focus true")
      }
    );
  }


  onSubmitTemplateBased(receiptFromForm: Receipt) { 
    console.log("this.receipt: ", this.receipt)
    console.log("receiptFromForm: ", receiptFromForm)
    // TODO: UPDATE RECEIPT
    this.receiptDTO.id = this.receipt.id
    this.receiptDTO.model_year = receiptFromForm.model_year
    this.receiptDTO.make = receiptFromForm.make
    this.receiptDTO.vin = receiptFromForm.vin
    this.receiptDTO.first_name = receiptFromForm.first_name
    this.receiptDTO.last_name = receiptFromForm.last_name
    this.receiptDTO.state = receiptFromForm.state
    this.receiptDTO.date_added.String = this.receipt.date_added
    this.receiptDTO.date_added.Valid = true
    this.receiptDTO.qa_score.Int64 = receiptFromForm.qa_score
    this.receiptDTO.qa_score.Valid = false
    this.receiptDTO.qa_date.String = this.receipt.qa_date
    this.receiptDTO.qa_date.Valid = false
    console.log("receiptDTO: ", this.receiptDTO)
    this.receiptService.updateVerifiedReceipt(this.receiptDTO)
      .subscribe(res => {
        console.log("receipt updated successfully: ", res);
      });
    console.log("Emitting input event.");
    Emitters.inputEmitter.emit();
  }

  updateReceipt() {
    console.log("receipt update started")
  }

  refreshData() {
    this.http.get(this.getLatestUnverifiedReceiptURL, {withCredentials: true}).subscribe(
      (res: any) => {
        this.receipt.id = res.id;
        this.receipt.model_year = res.model_year;
        this.receipt.make = res.make;
        this.receipt.vin = res.vin;
        this.receipt.first_name = res.first_name;
        this.receipt.last_name = res.last_name;
        this.receipt.state = res.state;
        this.receipt.date_added = res.date_added.String;
        this.receipt.qa_score = res.qa_score.Int64;
        this.receipt.qa_date = res.qa_date.String;
        console.log("got unverified receipt:", res.id);
        this.modelYear.nativeElement.focus()

      },
      err => {
        console.error(err);        
      }
    );
  }
}
