# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Google SHAKEN cert 969H

Tested At: 22 Aug 24 15:41 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -637 day(s)\
Subject: O=Google, C=US, CN=Google SHAKEN cert 969H\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://www.gstatic.com/gtp/stir/sFCvi_7snpTdPyz7VUUx8g.pem

[View certificate details](https://x509.io/?cert=MIICxjCCAkugAwIBAgIJAKoe4K1JZ1j3MAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMjEwMjUwMTA4MzFaFw0yMjExMjQwMTA4MzFaMEAxIDAeBgNVBAMMF0dvb2dsZSBTSEFLRU4gY2VydCA5NjlIMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGR29vZ2xlMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJBcUMLKgZucgu6mk60Rp5Vy7rtAT7LKt4%2BSpX0P4rDC8s%2Bva6PxQX1yecf1163%2Bd3mT4nvTi%2Fri%2FSbq35edRnKOB3jCB2zAWBggrBgEFBQcBGgQKMAigBhYEOTY5SDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAfBgNVHSMEGDAWgBRxL8iC3KjgIuPfoGj5%2BF5chN7lvTAdBgNVHQ4EFgQUV5SQnhscZUB8hgi6PX%2FVBOtirdcwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNpADBmAjEAlsLIpfywDHGLW0Z06o6RDbvqAQhP6485S9vw%2BG0iAemkzuwA8oWoYJX6zI9AT8EqAjEA4koIb5jsGGRl%2FZBGXupQEN7QGoYcoaJjgbcsVhYWx4XR6PegUaLeW3RJhS3ZdQ9I)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 969H', but common name is 'Google SHAKEN cert 969H' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Aug 24 15:44 UTC