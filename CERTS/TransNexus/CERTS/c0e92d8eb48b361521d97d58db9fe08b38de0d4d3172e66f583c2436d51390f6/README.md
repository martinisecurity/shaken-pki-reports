# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 012K

Tested At: 04 Oct 24 15:32 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -711 day(s)\
Subject: CN=SHAKEN 012K, OU=SHAKEN, O=CallCurrent\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/06ebd24a-1f0a-46d5-8a2f-a7ae49be8ed6/1113803cceb280f3ccac101750f2ef94.pem

[View certificate details](https://x509.io/?cert=MIIC8jCCApmgAwIBAgIQeJRKs0QAAzc3xRG7bxWOlzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTcwODM4NDdaFw0yMjEwMjQwODM4NDZaMFAxCzAJBgNVBAYTAlVTMRowGAYDVQQKExFDYWxsQ3VycmVudCwgSW5jLjEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDEySzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOu9klx7D1kEJDjjytdQ6t9PEeXOScefJ05alnIV0MZgXWos2Lo5Q2LJWHskUzmVadhEWH1SFsrPWwpZpTi3y0ujggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUgcW49ksYA3UOgpZ1NNApUTU%2F7BQwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDEySzAKBggqhkjOPQQDAgNHADBEAiBwRu3tw6oqNtkg4CWH7iPd1tNXZikXDR%2F%2BEdvwS6Qm2wIgUXqJVKctuP%2FuMMzaaJigAojaRwseAjHhaQ6Z1EDw16s%3D)

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


Generated: 04 Oct 24 16:29 UTC