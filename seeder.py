from dotenv import load_dotenv
import psycopg2
from psycopg2.extensions import ISOLATION_LEVEL_AUTOCOMMIT
import os
from pathlib import Path
import csv
import re
import wget
import tarfile

url = 'https://s3.us-west-1.amazonaws.com/packform.documents/recruiting/test_data.tar.gz'

load_dotenv()

path = Path(__file__).parent

path = os.path.join(path, 'data')
 
if not os.path.exists(path):
    os.makedirs(path)
    wget.download(url,out = path)
    my_tar = tarfile.open(os.path.join(path,'test_data.tar.gz'))
    my_tar.extractall(os.path.join(path)) # specify which folder to extract to
    my_tar.close()




PORT = os.getenv("PORT")
PASSWORD = os.getenv("PASSWORD")
USER = os.getenv("USER")
DBNAME = os.getenv("DBNAME")

# Connect to PostgreSQL DBMS

con = psycopg2.connect("user="+USER+" password='{}'".format(PASSWORD))
con.set_isolation_level(ISOLATION_LEVEL_AUTOCOMMIT)

# Obtain a DB Cursor
cursor = con.cursor()

# Create DataBase statement

try:
  cursor.execute("CREATE DATABASE "+DBNAME+"")
except psycopg2.Error as e:
  print(e)

cursor.close()

def get_table_name(file_name):
  table_name = re.search('([a-z]+\.csv)', file_name).group(1)
  table_name = table_name.split('.')[0]
  return table_name


# Create a table in PostgreSQL database and dump the data from the csv files into the tables
def read_csv_and_insert_data():
  conn_table = psycopg2.connect("dbname="+DBNAME+" user= "+USER+" password= "+PASSWORD+"")
  cursor_table = conn_table.cursor()
  print("Connecting to Database")
  csv_files_list = os.listdir(path)
  for csv_file in csv_files_list:
    if 'gz' in csv_file:
      continue
    table_name = get_table_name(csv_file)
    fileInput = open(os.path.join(path, csv_file), "r")
    firstLine = fileInput.readline().strip()
    columns = firstLine.split(",")
    sqlQueryCreate = 'DROP TABLE IF EXISTS '+ table_name + ";\n"
    sqlQueryCreate += 'CREATE TABLE '+ table_name + "("
    #some loop or function according to your requiremennt
    # Define columns for table
    for column in columns:
      if column == 'id' or column == 'order_id' or column == 'company_id' or column == 'order_item_id':
        sqlQueryCreate += column + " INTEGER,\n"
      elif column == 'created_at':
        sqlQueryCreate += column + " TimeStamp,\n"
      else:
        sqlQueryCreate += column + " VARCHAR(64),\n"
    sqlQueryCreate = sqlQueryCreate[:-2]
    sqlQueryCreate += ");"

    try:
      result = cursor_table.execute(sqlQueryCreate)
    except psycopg2.Error as e:
      print(e)
    with open(os.path.join(path, csv_file)) as f:
      table_name = get_table_name(csv_file)
      next(f)
      spamreader = csv.reader(f, delimiter=',')
      for row in spamreader:
        print('.....', row)
        values = ('\', \''.join(row))
        values = values.replace(',,', ',')
        insert_sql_statement = "Insert into "+table_name+" ({0}) values ('{1}')".format(",".join(columns), values)
        print(insert_sql_statement)
        cursor_table.execute(insert_sql_statement)


  conn_table.commit()
  cursor_table.close()       


read_csv_and_insert_data()



