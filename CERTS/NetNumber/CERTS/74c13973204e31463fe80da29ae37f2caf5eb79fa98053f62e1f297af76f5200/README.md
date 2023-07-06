# STIR/SHAKEN CA Ecosystem Compliance

## Certificate EssexTel INC SHAKEN 692J

Tested At: 06 Jul 23 14:06 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 3 day(s)\
Subject: O=EssexTel INC SHAKEN 692J, CN=EssexTel INC SHAKEN 692J\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://netnumber-sti-cr.s3.amazonaws.com/certs/041091ff-fdf3-400e-b212-fe7f448439bd

[View certificate details](https://understandingwebpki.com/?cert=MIICyjCCAlGgAwIBAgIJAMqjANW13SrkMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMzA2MDkwOTE4MzRaFw0yMzA3MDkwOTE4MzRaMEYxITAfBgNVBAMMGEVzc2V4VGVsIElOQyBTSEFLRU4gNjkySjEhMB8GA1UECgwYRXNzZXhUZWwgSU5DIFNIQUtFTiA2OTJKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEOSKR3RwF1Jxa0qkpWt8GcyiloYEwtarSg%2FJ5%2BB4gkzDi2Amfbb5%2F3%2BuAQt94vwb7VlTyPfRZ5s7oo27mP7M8dqOB3jCB2zAWBggrBgEFBQcBGgQKMAigBhYENjkySjAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAfBgNVHSMEGDAWgBRxL8iC3KjgIuPfoGj5%2BF5chN7lvTAdBgNVHQ4EFgQUAgSwG13nn%2BBaMkP%2FlA%2Bys9DBXZMwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNnADBkAjBZmL%2Bj9r13eTxgLQosdUN71dMJYK31acoFDXGNXpN3EFoQRC9i6se6rVEBynKrDh0CMFdAYEfERZk6aQ9iy4ij84Dd1NHyXuib61xzgw9MH2IdfqnosY6UKOzmDKMT28FFVg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject](../../ISSUES/e_atis_subject/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |


Generated: 06 Jul 23 14:08 UTC