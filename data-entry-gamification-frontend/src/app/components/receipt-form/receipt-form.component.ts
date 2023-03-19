import { Component, OnInit, ElementRef, ViewChild } from '@angular/core';
import { Receipt } from 'src/app/entities/receipt';
import { ReceiptService } from 'src/app/services/receipt.service';
import { Emitters } from 'src/app/emitters/emitters';
import {DatePipe, formatDate} from '@angular/common';

@Component({
  selector: 'app-receipt-form',
  templateUrl: './receipt-form.component.html',
  styleUrls: ['./receipt-form.component.css']
})
export class ReceiptFormComponent implements OnInit {
  @ViewChild("modelYear", {  }) modelYear: ElementRef;

  receipt: Receipt = {id: 0, model_year: 0, make: "", vin: "", first_name: "", last_name: "", state: "", date_added: "", qa_score: 0, qa_date: ""};

  constructor(
    private receiptService: ReceiptService
  ) {}

  ngOnInit(): void {
    Emitters.inputEmitter.subscribe(
      () => {
        this.modelYear.nativeElement.focus()
      }
    );
  }
  

  onSubmitTemplateBased(receiptFromForm: Receipt) { 
    receiptFromForm.id = 0;
    if (this.receipt.date_added === "") {
      // this.receipt.date_added = new Date().toISOString(); // UTC
      var g = new Date()
      var eg = formatDate(g, 'yyyy-MM-dd hh:mm:ss Z UTC', "en-US", "UTC")
      console.log("eg " + eg)
      // this.receipt.date_added = this.datePipe.transform(new Date(), 'yyyy-MM-dd hh:mm:ss Z', "UTC");
      this.receipt.date_added = eg.toString();
    }
    receiptFromForm.date_added = this.receipt.date_added;
    this.receiptService.addReceipt(receiptFromForm)
      .subscribe(receiptFromForm => {
        console.log("receipt added successfully: ", receiptFromForm);
      });
    Emitters.inputEmitter.emit();
  }
  
  newReceipt() {
    console.log("new receipt added")
  }

}
