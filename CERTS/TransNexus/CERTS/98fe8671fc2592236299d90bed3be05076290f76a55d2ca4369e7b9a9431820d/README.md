# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 721J

Tested At: 04 Oct 24 15:42 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -722 day(s)\
Subject: CN=SHAKEN 721J, OU=SHAKEN, O=True IP Solutions, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/a1df4bf4-8858-47bb-9388-835c5c7cb5c4/c444a025b79edfc6845719a4946375dd.pem

[View certificate details](https://x509.io/?cert=MIIC8zCCApmgAwIBAgIQSuf6PB%2FuBptTk2zpxTQCNzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDUyMDIwMTBaFw0yMjEwMTIyMDIwMDlaMFAxCzAJBgNVBAYTAlVTMRowGAYDVQQKExFUcnVlIElQIFNvbHV0aW9uczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNzIxSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABH1JQi2MSxvoPv34ll80EDc%2B6OO7uSwtQBOGsPOXijEtL405UEyoHVAouzQc0rMmWw4fEwT2wX6QkbjlBUuyoCajggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUWtGIhoFzLHl5zkz0lZMsUyMNOfowHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzIxSjAKBggqhkjOPQQDAgNIADBFAiBtml4RHIEs5y3A6oCxwq8dxJfN0rcLKBqpxyloBtd5KQIhAOjubK13W9rKtTIXHXzBGa97Vjp%2F487rvlPgWDxpjvKG)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC