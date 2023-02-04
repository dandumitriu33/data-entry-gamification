import { Component } from '@angular/core';
import { Receipt } from 'src/app/entities/receipt';
import { ReceiptService } from 'src/app/services/receipt.service';
import { ActivatedRoute } from '@angular/router';
import { Emitters } from 'src/app/emitters/emitters';

@Component({
  selector: 'app-receipt-form',
  templateUrl: './receipt-form.component.html',
  styleUrls: ['./receipt-form.component.css']
})
export class ReceiptFormComponent {

  constructor(
    private receiptService: ReceiptService,
    private route: ActivatedRoute
  ) {}

  receipt: Receipt = {id: 0, model_year: 0, make: "", vin: "", first_name: "", last_name: "", state: ""};

  onSubmitTemplateBased(receiptFromForm: Receipt) { 
    console.log("this.receipt: ", this.receipt)
    receiptFromForm.id = 0;
    console.log("receiptFromForm: ", receiptFromForm)
    this.receiptService.addReceipt(receiptFromForm)
      .subscribe(receiptFromForm => {
        console.log("receipt added successfully: ", receiptFromForm);
      });
    console.log("Emitting input event.");
    let dateTime = new Date();
    Emitters.inputEmitter.emit(dateTime);
  }

  newReceipt() {
    console.log("new receipt added")
  }

}
