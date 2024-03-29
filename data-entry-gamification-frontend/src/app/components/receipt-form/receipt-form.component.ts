import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { Receipt } from 'src/app/entities/receipt';
import { ReceiptService } from 'src/app/services/receipt.service';
import { Emitters } from 'src/app/emitters/emitters';
import { formatDate} from '@angular/common';
import { Validators, FormBuilder, ValidatorFn, AbstractControl, ValidationErrors } from '@angular/forms';

@Component({
  selector: 'app-receipt-form',
  templateUrl: './receipt-form.component.html',
  styleUrls: ['./receipt-form.component.css']
})
export class ReceiptFormComponent implements OnInit {
  @ViewChild('model_year_reactive_autofocus') modelYearReactiveElement: ElementRef;
  showSuccessMessage = false;

  receiptFormGroup = this.fb.group({
    model_year_reactive: ['', [Validators.required, Validators.min(1800), Validators.max(2200), Validators.pattern(/^\d{4}$/)]],
    make_reactive: ['', Validators.required],
    vin_reactive: ['', [Validators.required, Validators.minLength(17), Validators.maxLength(17)]],
    first_name_reactive: ['', [Validators.required, this.forbiddenNameValidator(/boob/i)]],
    last_name_reactive: ['', Validators.required],
    state_reactive: ['', Validators.required],
  })  

  constructor(
    private receiptService: ReceiptService,
    private fb: FormBuilder
  ) {}

  ngOnInit(): void {
    setTimeout(() => {
      this.focusInputElement();
    }, 100);
  }

  get model_year_reactive() { return this.receiptFormGroup.get('model_year_reactive'); }
  get make_reactive() { return this.receiptFormGroup.get('make_reactive'); }
  get vin_reactive() { return this.receiptFormGroup.get('vin_reactive'); }
  get first_name_reactive() { return this.receiptFormGroup.get('first_name_reactive'); }
  get last_name_reactive() { return this.receiptFormGroup.get('last_name_reactive'); }
  get state_reactive() { return this.receiptFormGroup.get('state_reactive'); }  

  onSubmit(){
    let tempId = 0
    let tempDate = formatDate(new Date(), 'yyyy-MM-dd hh:mm:ss Z UTC', "en-US", "UTC").toString()
    let receiptFromForm = new Receipt(tempId, 0, "", "", "", "", "", tempDate, 0, "");
    
    receiptFromForm.model_year = Number(this.receiptFormGroup.controls['model_year_reactive'].value??"0")
    receiptFromForm.make = this.receiptFormGroup.controls['make_reactive'].value??"N/A"
    receiptFromForm.vin = this.receiptFormGroup.controls['vin_reactive'].value??"N/A"
    receiptFromForm.first_name = this.receiptFormGroup.controls['first_name_reactive'].value??"N/A"
    receiptFromForm.last_name = this.receiptFormGroup.controls['last_name_reactive'].value??"N/A"
    receiptFromForm.state = this.receiptFormGroup.controls['state_reactive'].value??"N/A"
    this.receiptService.addReceipt(receiptFromForm)
      .subscribe(receiptFromForm => {
        this.showSuccessMessage = true;
        console.log("Receipt from formGroup added successfully: ", receiptFromForm);
        this.receiptFormGroup.reset();
        setTimeout(() => {
          this.focusInputElement();
        }, 100);
        setTimeout(() => {
          this.showSuccessMessage = false;
        }, 2000);
      });
    Emitters.inputEmitter.emit();
  }

  /** A First Name can't match the given regular expression */
  forbiddenNameValidator(nameRe: RegExp): ValidatorFn {
    return (control: AbstractControl): ValidationErrors | null => {
      const forbidden = nameRe.test(control.value);
      return forbidden ? {forbiddenName: {value: control.value}} : null;
    };
  }

  focusInputElement() {
    if (this.modelYearReactiveElement && this.modelYearReactiveElement.nativeElement) {
      this.modelYearReactiveElement.nativeElement.focus();        
    } 
  }

}
