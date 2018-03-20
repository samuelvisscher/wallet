import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { KittenDetailComponent } from './kitten-detail.component';

describe('KittenDetailComponent', () => {
  let component: KittenDetailComponent;
  let fixture: ComponentFixture<KittenDetailComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ KittenDetailComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KittenDetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
