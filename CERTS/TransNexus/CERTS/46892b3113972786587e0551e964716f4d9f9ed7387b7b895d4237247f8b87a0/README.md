# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 625J

Tested At: 04 Oct 24 15:43 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -704 day(s)\
Subject: CN=SHAKEN 625J, OU=SHAKEN, O=Victory Telecom Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/b7343686-5ed8-402c-89a3-8bf1a3d48975/9152c149094a3c0a73bbaf2924b00f92.pem

[View certificate details](https://x509.io/?cert=MIIC9zCCApygAwIBAgIQWag7ekSU1xAPxclpt%2BwJHjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMyMDIyNTdaFw0yMjEwMzAyMDIyNTZaMFMxCzAJBgNVBAYTAlVTMR0wGwYDVQQKExRWaWN0b3J5IFRlbGVjb20gSW5jLjEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNjI1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNiOHfTz1ERRunK5jHWdGxDRE14JO56leMcwxP31KvTvJJ6nd4CmFT9nt8vTZdcOW%2BKeuaomiuRA2uSnLY08ZmujggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUl9zk0W3epYG41L50uHHibNRnweowHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENjI1SjAKBggqhkjOPQQDAgNJADBGAiEAyQLoNFAaD7%2Bp7VHG6l1gv5AStCKh%2FChsi95Ri%2BvfrNMCIQCpE9H13yTO7gnMDAKWBJy4TUOa1uXbK7ZVbzW84T3%2FWg%3D%3D)

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