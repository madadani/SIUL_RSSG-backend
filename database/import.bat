@echo off
set PGPASSWORD=root
psql -h localhost -p 5432 -U postgres -d siul_db -f schema_full.sql
pause
