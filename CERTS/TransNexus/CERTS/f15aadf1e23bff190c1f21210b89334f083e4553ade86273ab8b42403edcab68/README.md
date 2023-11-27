# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 089K

Tested At: 27 Nov 23 23:13 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -411 day(s)\
Subject: CN=SHAKEN 089K, OU=SHAKEN, O=Logista Solutions, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/5f7c135d-caf3-4661-abb1-d1720e7872e2/1022ee180b7db2f1ca2ec253c94dfad0.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9DCCApmgAwIBAgIQW%2BolII37fPiusPpNDa%2BqUDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDUxODM0NDJaFw0yMjEwMTIxODM0NDFaMFAxCzAJBgNVBAYTAlVTMRowGAYDVQQKExFMb2dpc3RhIFNvbHV0aW9uczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDg5SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABG5ucG%2BFq6GmSm16oGufFhZpL12srCP6J4Nd%2FuVXlHWaXJFe4k9ymi2nirIj6XTvzHCP%2FqMTRfOnDApJXI47qoWjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUVgg2%2BmiopE23swW1p1cc7Cp%2Fq78wHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDg5SzAKBggqhkjOPQQDAgNJADBGAiEAt3H1LUM3U1xfDjCr2%2BHzVtpy95mpFGHK6xMA%2FwNNsfMCIQDmRoZc%2FETe%2FkFmM%2Bl38JQMpJ2ErN%2Fep08bchNqTosmFA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 27 Nov 23 23:28 UTC