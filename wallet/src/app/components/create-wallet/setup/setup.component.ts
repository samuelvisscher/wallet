import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-setup',
  templateUrl: './setup.component.html',
  styleUrls: ['./setup.component.scss']
})
export class SetupComponent implements OnInit {

  @Output() onSubmit = new EventEmitter();

  form: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
  ) { }

  ngOnInit() {
    this.form = this.formBuilder.group({
      name: ['', Validators.required],
    });
  }

  submit() {
    this.onSubmit.emit(this.form.value);
  }
}
