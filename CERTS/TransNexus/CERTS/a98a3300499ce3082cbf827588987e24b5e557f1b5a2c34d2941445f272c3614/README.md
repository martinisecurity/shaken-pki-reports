# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 459J

Tested At: 04 Oct 24 15:45 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -725 day(s)\
Subject: CN=SHAKEN 459J, OU=SHAKEN, O=Altaworx, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/d3d6dbfc-2914-49c7-8f47-d0aa5bd5d764/b9ac4c13c48a2c103d91a1737e93c52a.pem

[View certificate details](https://x509.io/?cert=MIIC6jCCApCgAwIBAgIQSx4GCBPgNqsIeJbUsBKtOjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDIyMDIxMDBaFw0yMjEwMDkyMDIwNTlaMEcxCzAJBgNVBAYTAlVTMREwDwYDVQQKEwhBbHRhd29yeDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNDU5SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCHeEB%2FNgrCJ4jivWz7Q%2BEgrtQfQ%2BJicShC6NCmHKt0XDQY0JcxYcMyKuEV0aEcq7CIT5n3UgpQF%2FD8ii0SuNWyjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUJsz9lFz6EWSHk3v94Ng%2FH98xVTswHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENDU5SjAKBggqhkjOPQQDAgNIADBFAiBtHdNgqMTa5tpatlc76DMmj4D9jiA8XteY30J8Rs%2F%2FegIhAJUDtCFjd4ZG5JUns1uuiDaT7BpOHvknAdxX220s44Br)

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