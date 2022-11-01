# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

### Leaf Certificates

- 771 certificates were included in the corpus being tested
- 111 certificates in the corpus were skipped because they are duplicates
- 307 certificates in the corpus were skipped because they are expired
- 2 certificates in the corpus were skipped because they are not currently trusted
- 351 certificates being tested against the remaining rules
- 96.30% of certificates contain one or more Error level issue
- 75.50% of certificates contain one or more Warning level issue
- 1.99% of certificates contain one or more Notice level issue
- 15.38% of certificates are too old to be assessed against currently enforced expectations
- 284 days is the average remaining validity for the certificates in the corpus
- 284 days is the average initial validity for the certificates in the corpus
- 181 certificates expire in the next 30 days

### CA Certificates

- 33 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 2 certificates in the corpus were skipped because they are expired
- 2 certificates in the corpus were skipped because they are not currently trusted
- 29 certificates being tested against the remaining rules
- 55.17% of certificates contain one or more Error level issue
- 86.21% of certificates contain one or more Warning level issue
- 3.45% of certificates contain one or more Notice level issue
- 75.86% of certificates are too old to be assessed against currently enforced expectations
- 5741 days is the average remaining validity for the certificates in the corpus
- 5630 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

## Certificate Repository

- 51.57% of certificate repositories contain one or more Error level issue
- 77.21% of certificates repositories contain one or more Warning level issue
- 0.00% of certificates repositories contain one or more Notice level issue

## Details

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 45 (12.82%) | 45 (12.82%) | 45 (12.82%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 1 (0.28%) | 1 (0.28%) | 1 (0.28%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 7 (1.99%) | 0 (0.00%) | 0 (0.00%) | 4 (1.14%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 21 (5.98%) | 21 (5.98%) | 6 (1.71%) | 0 (0.00%) | 20 (5.70%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 3 (0.85%) | 3 (0.85%) | 1 (0.28%) | 3 (0.85%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 95 (27.07%) | 89 (25.36%) | 33 (9.40%) | 0 (0.00%) | 32 (9.12%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 8 (2.28%) | 8 (2.28%) | 8 (2.28%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 3 (0.85%) | 3 (0.85%) | 3 (0.85%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 95 (27.07%) | 95 (27.07%) | 95 (27.07%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 1 (0.28%) | 1 (0.28%) | 1 (0.28%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 72 (20.51%) | 72 (20.51%) | 72 (20.51%) | 0 (0.00%) | 2 (0.57%) |
| **Total** | 351 (100.00%) | 338 (96.30%) | 265 (75.50%) | 7 (1.99%) | 54 (15.38%) |

### CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 2 (6.90%) | 0 (0.00%) | 2 (6.90%) |
| [GBSDTech](CERTS/GBSDTech/README.md#ca-certificates) | 2 (6.90%) | 1 (3.45%) | 2 (6.90%) | 0 (0.00%) | 2 (6.90%) |
| [Martini Security](CERTS/Martini_Security/README.md#ca-certificates) | 3 (10.34%) | 0 (0.00%) | 3 (10.34%) | 1 (3.45%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#ca-certificates) | 2 (6.90%) | 2 (6.90%) | 0 (0.00%) | 0 (0.00%) | 2 (6.90%) |
| [NetNumber](CERTS/NetNumber/README.md#ca-certificates) | 3 (10.34%) | 2 (6.90%) | 3 (10.34%) | 0 (0.00%) | 3 (10.34%) |
| [Neustar](CERTS/Neustar/README.md#ca-certificates) | 6 (20.69%) | 4 (13.79%) | 6 (20.69%) | 0 (0.00%) | 4 (13.79%) |
| [Peeringhub](CERTS/Peeringhub/README.md#ca-certificates) | 2 (6.90%) | 2 (6.90%) | 2 (6.90%) | 0 (0.00%) | 1 (3.45%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#ca-certificates) | 2 (6.90%) | 2 (6.90%) | 2 (6.90%) | 0 (0.00%) | 2 (6.90%) |
| [Sansay](CERTS/Sansay/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 2 (6.90%) | 0 (0.00%) | 1 (3.45%) |
| [T-Mobile](CERTS/T-Mobile/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (6.90%) |
| [TransNexus](CERTS/TransNexus/README.md#ca-certificates) | 3 (10.34%) | 3 (10.34%) | 3 (10.34%) | 0 (0.00%) | 3 (10.34%) |
| **Total** | 29 (100.00%) | 16 (55.17%) | 25 (86.21%) | 1 (3.45%) | 22 (75.86%) |

### Key

| Type | Description |
|------|-------------|
| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notices | Tests in which industry best practices are not followed. |
| Not Effective | Tests that exist in the current specifications but were not in effect at the time of issuance. |


Generated: 01/11/2022 at 20:34:21