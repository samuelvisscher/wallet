import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ShowSeedComponent } from './show-seed.component';

describe('ShowSeedComponent', () => {
  let component: ShowSeedComponent;
  let fixture: ComponentFixture<ShowSeedComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ShowSeedComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ShowSeedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
