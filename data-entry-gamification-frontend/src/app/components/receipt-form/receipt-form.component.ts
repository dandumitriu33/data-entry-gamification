import { Component, OnInit, ElementRef, ViewChild } from '@angular/core';
import { Receipt } from 'src/app/entities/receipt';
import { ReceiptService } from 'src/app/services/receipt.service';
import { ActivatedRoute } from '@angular/router';
import { Emitters } from 'src/app/emitters/emitters';

@Component({
  selector: 'app-receipt-form',
  templateUrl: './receipt-form.component.html',
  styleUrls: ['./receipt-form.component.css']
})
export class ReceiptFormComponent implements OnInit {
  @ViewChild("modelYear", {  }) modelYear: ElementRef;

  receipt: Receipt = {id: 0, model_year: 0, make: "", vin: "", first_name: "", last_name: "", state: "", date_added: "", qa_score: 0, qa_date: ""};

  constructor(
    private receiptService: ReceiptService,
    private route: ActivatedRoute
  ) {}

  ngOnInit(): void {
    Emitters.inputEmitter.subscribe(
      () => {
        this.modelYear.nativeElement.focus()
        console.log("focus true")
      }
    );
  }
  

  onSubmitTemplateBased(receiptFromForm: Receipt) { 
    console.log("this.receipt: ", this.receipt)
    receiptFromForm.id = 0;
    console.log("receiptFromForm: ", receiptFromForm)
    this.receiptService.addReceipt(receiptFromForm)
      .subscribe(receiptFromForm => {
        console.log("receipt added successfully: ", receiptFromForm);
      });
    console.log("Emitting input event.");
    Emitters.inputEmitter.emit();
  }

  onSubmitTemplateBased2(receiptFromForm: Receipt) { 
    console.log("this.receipt: ", this.receipt)
    receiptFromForm.id = 0;
    console.log("receiptFromForm: ", receiptFromForm)
    this.receiptService.addReceipt(receiptFromForm)
      .subscribe(receiptFromForm => {
        console.log("receipt added successfully: ", receiptFromForm);
      });
    console.log("Emitting input event.");
    Emitters.inputEmitter.emit();
  }

  newReceipt() {
    console.log("new receipt added")
  }

}
