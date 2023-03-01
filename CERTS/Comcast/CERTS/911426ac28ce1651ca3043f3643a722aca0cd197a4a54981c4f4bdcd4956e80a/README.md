# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 01 Mar 23 18:22 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 13 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/7b16351a-4880-44e6-a0d5-442255d92316.cer

[View certificate details](https://understandingwebpki.com/?cert=MIICVTCCAfygAwIBAgIHPmMGVjsm3jAKBggqhkjOPQQDAjBfMQswCQYDVQQGEwJVUzEVMBMGA1UECAwMUGVubnN5bHZhbmlhMRAwDgYDVQQKDAdDb21jYXN0MScwJQYDVQQDDB5Db21jYXN0IFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EwHhcNMjMwMjEyMTExMjI0WhcNMjMwMzE0MTExMjI0WjBeMQswCQYDVQQGEwJVUzEVMBMGA1UECBMMUGVubnN5bHZhbmlhMRUwEwYDVQQHEwxQaGlsYWRlbHBoaWExEDAOBgNVBAoTB0NvbWNhc3QxDzANBgNVBAMTBlNIQUtFTjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABF1tOceInzhgPzuJY65ORXSzsZDznYhUwxnIT9MsvbAsXMzDdSYGj2I99tfH33EdIBFybhD5K1u0%2F0AR9L3sptCjgaMwgaAwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUv44cYFUkxRFIz0JFRL%2Bw9Dor5igwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBYGCCsGAQUFBwEaBAowCKAGFgQzMThKMAoGCCqGSM49BAMCA0cAMEQCIE5HmHluQqFoJsEPWaxw%2Fs8hmBqdXKQ8BG1dJxfIhtFrAiBgOB44rsfgGcFZLhUQlg87cAR3a578q0VhHEAQPqyJxQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_serial_number](../../ISSUES/e_atis_serial_number/README.md) | error | ATIS1000080 | STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG) |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain a Subject Key Identifier extension |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 318J' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 01 Mar 23 18:22 UTC