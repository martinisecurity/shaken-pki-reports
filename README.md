# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

337 certificates were included in the corpus being tested\
302 certificates in the corpus were skipped because they were expired\
108 repositories in the corpus were skipped because they were duplicated\
2 certificates in the corpus were skipped because they are not currently trusted\
93.77% of certificates contain one or more Error level issue\
73.59% of certificates contain one or more Warning level issue\
2.37% of certificates contain one or more Notice level issue\
16.32% of certificates are too old to be assessed against currently enforced expectations\
295 days is the average remaining validity for the certificates in the corpus\
295 days is the average initial validity for the certificates in the corpus\
168 certificates expire in the next 30 days

### CA Certificates

29 certificates were included in the corpus being tested\
2 certificates in the corpus were skipped because they were expired\
0 repositories in the corpus were skipped because they were duplicated\
2 certificates in the corpus were skipped because they are not currently trusted\
41.38% of certificates contain one or more Error level issue\
86.21% of certificates contain one or more Warning level issue\
62.07% of certificates contain one or more Notice level issue\
75.86% of certificates are too old to be assessed against currently enforced expectations\
5741 days is the average remaining validity for the certificates in the corpus\
5630 days is the average initial validity for the certificates in the corpus\
0 certificates expire in the next 30 days

## Certificate Repository

53.41% of certificate repositories contain one or more Error level issue\
78.93% of certificates repositories contain one or more Warning level issue\
0.00% of certificates repositories contain one or more Notice level issue

## Details

## Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 43 (12.76%) | 43 (12.76%) | 43 (12.76%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 1 (0.30%) | 1 (0.30%) | 1 (0.30%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 7 (2.08%) | 0 (0.00%) | 0 (0.00%) | 4 (1.19%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 21 (6.23%) | 21 (6.23%) | 6 (1.78%) | 0 (0.00%) | 20 (5.93%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 4 (1.19%) | 4 (1.19%) | 1 (0.30%) | 4 (1.19%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 96 (28.49%) | 82 (24.33%) | 32 (9.50%) | 0 (0.00%) | 33 (9.79%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 5 (1.48%) | 5 (1.48%) | 5 (1.48%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 3 (0.89%) | 3 (0.89%) | 3 (0.89%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 91 (27.00%) | 91 (27.00%) | 91 (27.00%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 1 (0.30%) | 1 (0.30%) | 1 (0.30%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 65 (19.29%) | 65 (19.29%) | 65 (19.29%) | 0 (0.00%) | 2 (0.59%) |
| **Total** | 337 (100.00%) | 316 (93.77%) | 248 (73.59%) | 8 (2.37%) | 55 (16.32%) |

## CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 2 (6.90%) | 1 (3.45%) | 2 (6.90%) |
| [GBSDTech](CERTS/GBSDTech/README.md#ca-certificates) | 2 (6.90%) | 1 (3.45%) | 2 (6.90%) | 0 (0.00%) | 2 (6.90%) |
| [Martini Security](CERTS/Martini_Security/README.md#ca-certificates) | 3 (10.34%) | 0 (0.00%) | 3 (10.34%) | 1 (3.45%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#ca-certificates) | 2 (6.90%) | 2 (6.90%) | 0 (0.00%) | 0 (0.00%) | 2 (6.90%) |
| [NetNumber](CERTS/NetNumber/README.md#ca-certificates) | 3 (10.34%) | 2 (6.90%) | 3 (10.34%) | 0 (0.00%) | 3 (10.34%) |
| [Neustar](CERTS/Neustar/README.md#ca-certificates) | 6 (20.69%) | 3 (10.34%) | 6 (20.69%) | 6 (20.69%) | 4 (13.79%) |
| [Peeringhub](CERTS/Peeringhub/README.md#ca-certificates) | 2 (6.90%) | 1 (3.45%) | 2 (6.90%) | 2 (6.90%) | 1 (3.45%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 2 (6.90%) | 2 (6.90%) | 2 (6.90%) |
| [Sansay](CERTS/Sansay/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 2 (6.90%) | 1 (3.45%) | 1 (3.45%) |
| [T-Mobile](CERTS/T-Mobile/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 0 (0.00%) | 2 (6.90%) | 2 (6.90%) |
| [TransNexus](CERTS/TransNexus/README.md#ca-certificates) | 3 (10.34%) | 3 (10.34%) | 3 (10.34%) | 3 (10.34%) | 3 (10.34%) |
| **Total** | 29 (100.00%) | 12 (41.38%) | 25 (86.21%) | 18 (62.07%) | 22 (75.86%) |


Generated: 31/10/2022 at 20:32:42