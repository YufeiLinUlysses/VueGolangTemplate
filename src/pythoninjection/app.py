import pymongo #pacakge connects to mongodb

#Connect to the mongodb portal
myclient = pymongo.MongoClient("mongodb://localhost:27017/")
#Get the particular database in mongodb
mydb = myclient["thepolyglotdeveloper"]
#Get the particular collection in the database
mycol = mydb["people"]

#Register a series of json data
mylist = [
  {"firstname":"Tom","lastname":"Jerry"},
  {"firstname":"Jerry","lastname":"Tom"}
]

#Insert data to the database
x = mycol.insert_many(mylist)

#print list of the _id values of the inserted documents:
print(x.inserted_ids) 