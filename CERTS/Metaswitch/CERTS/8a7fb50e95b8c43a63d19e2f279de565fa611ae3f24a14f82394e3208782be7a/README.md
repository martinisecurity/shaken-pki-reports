# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Metaswitch STI-CA SHAKEN Issuing 1

Tested At: 21 Nov 23 17:52 UTC\
Initial Validity Period: 4380 day(s)\
Remaining Validity Period: 4096 day(s)\
Subject: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Issuer: CN=Metaswitch STI-CA SHAKEN Root

[View certificate details](https://understandingwebpki.com/?cert=MIICfDCCAiKgAwIBAgIQaQRv6d8JnAo38%2Ffj5Kf%2F%2BjAKBggqhkjOPQQDAjAoMSYwJAYDVQQDDB1NZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gUm9vdDAeFw0yMzAyMTAxNDM4NDZaFw0zNTAyMDcxNDM4NDZaMC0xKzApBgNVBAMMIk1ldGFzd2l0Y2ggU1RJLUNBIFNIQUtFTiBJc3N1aW5nIDEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT8rXFLu6%2B6yRdl8bcBkaYj8zn05JVmeRYW1wPXttNOhQWK%2BbYC0lE5tR3IpgMou0xcLoubgVxr4kIsVo1UmMNZo4IBJzCCASMwHQYDVR0OBBYEFM0epwAQENoyHWkaOdXSRgssPIfWMB8GA1UdIwQYMBaAFDlUaboMbPafV9GbATFFiqD7%2F3hqMA8GA1UdEwEB%2FwQFMAMBAf8wDgYDVR0PAQH%2FBAQDAgIEMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwCgYIKoZIzj0EAwIDSAAwRQIgTBcPq%2FMRU%2FWsgH7ILFQbwBZINU1IjDiMeow4caGDW1ACIQDMR4jNXwc1nsFZtAjJIvT%2FhGCTMl4UQ3cDc%2BGYX0RNGQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id_ca](../../ISSUES/e_shaken_certificate_policies_id_ca/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: [2.16.840.1.114569.1.1.3 2.16.840.1.114569.1.1.4] |
| [e_atis_ca_subject_o_required](../../ISSUES/e_atis_ca_subject_o_required/README.md) | error | ATIS1000080 | The DN does not contain exactly one Organization (O=) attribute. |
| [e_atis_ca_subject_c_iso](../../ISSUES/e_atis_ca_subject_c_iso/README.md) | error | ATIS1000080 | Subject MUST be present and MUST contain exactly one value for Country (C=). |
| [e_atis_ca_subject_dn](../../ISSUES/e_atis_ca_subject_dn/README.md) | error | ATIS1000080 | Subject DN does not contain a Country (C=) attribute |


Generated: 21 Nov 23 17:53 UTC