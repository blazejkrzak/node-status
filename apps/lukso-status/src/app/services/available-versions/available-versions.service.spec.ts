import { TestBed } from '@angular/core/testing';

import { SoftwareService } from './available-versions.service';

describe('AvailableVersionsService', () => {
  let service: SoftwareService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SoftwareService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
