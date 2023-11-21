# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 089K

Tested At: 21 Nov 23 17:38 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -407 day(s)\
Subject: CN=SHAKEN 089K, OU=SHAKEN, O=Logista Solutions, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/5f7c135d-caf3-4661-abb1-d1720e7872e2/d5b8fef3be8fc769bfe4063c5ba64a2c.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApmgAwIBAgIQeJ4Tj2puGvcmhXZM2AZvGTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDIxODM0NDJaFw0yMjEwMDkxODM0NDFaMFAxCzAJBgNVBAYTAlVTMRowGAYDVQQKExFMb2dpc3RhIFNvbHV0aW9uczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDg5SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABEtjReStSR3nMXonbCbzr9xW0jwXdMsbo2Qh4qTUadIf9bmJD9EXeSmbLd%2BvLbQJK%2FzSbhOdQ7cvQqbteWffrV2jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUJOxqzTOqVWMCwQtoTf8%2FuQvK1nwwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDg5SzAKBggqhkjOPQQDAgNHADBEAiBkricXIO0rTYcfIE0iniF69ockSo8WapbWYUphW7WO3QIgdzsOX894R%2FD56pXl4aYAJyKNZ%2FJbaWZ53YFiPAZ%2F6EE%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:53 UTC