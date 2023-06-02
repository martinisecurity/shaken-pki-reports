# STIR/SHAKEN CA Ecosystem Compliance

## Certificate EssexTel INC SHAKEN 692J

Tested At: 02 Jun 23 01:10 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 7 day(s)\
Subject: O=EssexTel INC SHAKEN 692J, CN=EssexTel INC SHAKEN 692J\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://netnumber-sti-cr.s3.amazonaws.com/certs/419b6f62-c79c-404d-9253-36db8012193f

[View certificate details](https://understandingwebpki.com/?cert=MIICyjCCAlGgAwIBAgIJAPOMKzfYYIUXMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMzA1MTAwMDEwNTdaFw0yMzA2MDkwMDEwNTdaMEYxITAfBgNVBAMMGEVzc2V4VGVsIElOQyBTSEFLRU4gNjkySjEhMB8GA1UECgwYRXNzZXhUZWwgSU5DIFNIQUtFTiA2OTJKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEOSKR3RwF1Jxa0qkpWt8GcyiloYEwtarSg%2FJ5%2BB4gkzDi2Amfbb5%2F3%2BuAQt94vwb7VlTyPfRZ5s7oo27mP7M8dqOB3jCB2zAWBggrBgEFBQcBGgQKMAigBhYENjkySjAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAfBgNVHSMEGDAWgBRxL8iC3KjgIuPfoGj5%2BF5chN7lvTAdBgNVHQ4EFgQUAgSwG13nn%2BBaMkP%2FlA%2Bys9DBXZMwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNnADBkAjA3ihUyqnu17ynFSqfKG0erFya%2BjwJCZ2OWZ%2F4Gr52B6D23N6TrBathbk3QiC6RGQACMBZcJH0ueiAhvNhDdr80L6cO7cYHo5b1ihmdV5rI9RvJrJj7afebVUyL3bPyD5VKsg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_subject](../../ISSUES/e_atis_subject/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 02 Jun 23 01:12 UTC