#!/usr/bin/env bash
mongoimport --db bills --collection bills --drop --file user_bills.json
mongoimport --db bills --collection pu_catalog --drop --file pu_catalog.json
