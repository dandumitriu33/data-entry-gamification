import { Component, OnInit, ElementRef, ViewChild } from '@angular/core';
import { Receipt } from 'src/app/entities/receipt';
import { ReceiptService } from 'src/app/services/receipt.service';
import { Emitters } from 'src/app/emitters/emitters';
import { formatDate} from '@angular/common';
import { FormControl, FormGroup } from '@angular/forms';

@Component({
  selector: 'app-receipt-form',
  templateUrl: './receipt-form.component.html',
  styleUrls: ['./receipt-form.component.css']
})
export class ReceiptFormComponent implements OnInit {
  @ViewChild("modelYear", {  }) modelYear: ElementRef;

  receipt: Receipt = {id: 0, model_year: 0, make: "", vin: "", first_name: "", last_name: "", state: "", date_added: "", qa_score: 0, qa_date: ""};

  receiptFormGroup = new FormGroup({
    model_year_reactive: new FormControl(),
    make_reactive: new FormControl(''),
    vin_reactive: new FormControl(''),
    first_name_reactive: new FormControl(''),
    last_name_reactive: new FormControl(''),
    state_reactive: new FormControl(''),
  })
  

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
      var tempDate = formatDate(new Date(), 'yyyy-MM-dd hh:mm:ss Z UTC', "en-US", "UTC").toString()
      this.receipt.date_added = tempDate;
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

  onSubmit(){
    console.warn(this.receiptFormGroup.value);
    let tempId = 0
    let tempDate = formatDate(new Date(), 'yyyy-MM-dd hh:mm:ss Z UTC', "en-US", "UTC").toString()
    let receiptFromForm = new Receipt(tempId, 0, "", "", "", "", "", tempDate, 0, "");
    
    receiptFromForm.model_year = this.receiptFormGroup.controls['model_year_reactive'].value
    receiptFromForm.make = this.receiptFormGroup.controls['make_reactive'].value??"N/A"
    receiptFromForm.vin = this.receiptFormGroup.controls['vin_reactive'].value??"N/A"
    receiptFromForm.first_name = this.receiptFormGroup.controls['first_name_reactive'].value??"N/A"
    receiptFromForm.last_name = this.receiptFormGroup.controls['last_name_reactive'].value??"N/A"
    receiptFromForm.state = this.receiptFormGroup.controls['state_reactive'].value??"N/A"
    console.log(receiptFromForm)
    this.receiptService.addReceipt(receiptFromForm)
      .subscribe(receiptFromForm => {
        console.log("receipt from formGroup added successfully: ", receiptFromForm);
      });
    Emitters.inputEmitter.emit();
  }

}
