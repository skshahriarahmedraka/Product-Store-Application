## Product Store application (rest API) backend using Golang(Gin).

### For running the project:

spin up the Mongodb Database

```
cd deployment/docker
docker-compose up -d
```

run the project

```
go run main.go

```

### All the Routes

**Health**

```
"name": "health",
"request": {
   "method": "GET",
   "header": [],
   "url": {
      "raw": "localhost:8080/health",
      "host": [
         "localhost"
      ],
      "port": "8080",
      "path": [
         "health"
      ]
   }
},
"response": [
   {
      "name": "health",
      "originalRequest": {
         "method": "GET",
         "header": [],
         "url": {
            "raw": "localhost:8080/health",
            "host": [
               "localhost"
            ],
            "port": "8080",
            "path": [
               "health"
            ]
         }
      },
      "_postman_previewlanguage": null,
      "header": null,
      "cookie": [],
      "body": null
   }
]
```

**All Categories**

```
{
			"name": "All Catagories",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/category",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"category"
					]
				}
			},
			"response": [
				{
					"name": "All Catagories",
					"originalRequest": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/category",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"category"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

Database Table Schema Design:

brands(id, name, status_id, created_at)

```
categories(id, name, parent_id, sequence, status_id, created_at)

suppliers(id, name, email, phone, status_id, is_verified_supplier, created_at)

products(id, name, description, specifications, brand_id, category _id, supplier_id, unit_price,
discount_price, tags, status_id)

product_stocks(id, product_id, stock quantity, updated_at)
```

**Create Category**

```
{
			"name": "create category",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\", \n  \"name\": \"My\",\n  \"parent_id\": \"5f55987bfd86540813b7ec9b\",\n  \"sequence\": \"1\",\n  \"status_id\": true\n//   \"created_at\": \"2023-11-20T12:34:56Z\"\n}\n",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/category",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"category"
					]
				}
			},
			"response": [
				{
					"name": "create category",
					"originalRequest": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\", \n  \"name\": \"Example Category\",\n  \"parent_id\": \"5f55987bfd86540813b7ec9b\",\n  \"sequence\": \"1\",\n  \"status_id\": true\n//   \"created_at\": \"2023-11-20T12:34:56Z\"\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/category",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"category"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Get Category**

```
{
			"name": "get catagory",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/category/655b8131ff80fb1c35b10ac3",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"category",
						"655b8131ff80fb1c35b10ac3"
					]
				}
			},
			"response": [
				{
					"name": "get catagory",
					"originalRequest": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/category/655b8131ff80fb1c35b10ac3",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"category",
								"655b8131ff80fb1c35b10ac3"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Update Categry**

```
{
			"name": "Update Category",
			"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\", \n  \"name\": \"New Category\",\n  \"parent_id\": \"5f55987bfd86540813b7ec9b\",\n  \"sequence\": \"1\",\n  \"status_id\": true\n//   \"created_at\": \"2023-11-20T12:34:56Z\"\n}\n",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/category/655b8131ff80fb1c35b10ac3",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"category",
						"655b8131ff80fb1c35b10ac3"
					]
				}
			},
			"response": [
				{
					"name": "Update Category",
					"originalRequest": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\", \n  \"name\": \"New Category\",\n  \"parent_id\": \"5f55987bfd86540813b7ec9b\",\n  \"sequence\": \"1\",\n  \"status_id\": true\n//   \"created_at\": \"2023-11-20T12:34:56Z\"\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/category/655b8131ff80fb1c35b10ac3",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"category",
								"655b8131ff80fb1c35b10ac3"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

** Update Category**

```
{
			"name": "Delete Category",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "localhost:8080/category/655b8131ff80fb1c35b10ac3",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"category",
						"655b8131ff80fb1c35b10ac3"
					]
				}
			},
			"response": [
				{
					"name": "All Catagories Copy",
					"originalRequest": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "localhost:8080/category/655b8131ff80fb1c35b10ac3",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"category",
								"655b8131ff80fb1c35b10ac3"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Create Brand**

```
{
			"name": "create Brand",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\",\n  \"name\": \"Example Brand\",\n  \"status_id\": true\n//   \"created_at\": \"2023-11-20T12:34:56Z\" // replace with your actual timestamp value\n}\n",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/brand",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"brand"
					]
				}
			},
			"response": [
				{
					"name": "create Brand",
					"originalRequest": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\",\n  \"name\": \"Example Brand\",\n  \"status_id\": true\n//   \"created_at\": \"2023-11-20T12:34:56Z\" // replace with your actual timestamp value\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/brand",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"brand"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Get Brand**

```
{
			"name": "Get Brand",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/brand/655ba0168ed225db7c347301",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"brand",
						"655ba0168ed225db7c347301"
					]
				}
			},
			"response": [
				{
					"name": "Get Brand",
					"originalRequest": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/brand/655ba0168ed225db7c347301",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"brand",
								"655ba0168ed225db7c347301"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Update Brand**

```
{
			"name": "Update Brand",
			"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\",\n  \"name\": \"updated Brand\",\n  \"status_id\": true\n//   \"created_at\": \"2023-11-20T12:34:56Z\" // replace with your actual timestamp value\n}\n",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/brand/655ba0168ed225db7c347301",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"brand",
						"655ba0168ed225db7c347301"
					]
				}
			},
			"response": [
				{
					"name": "Update Brand",
					"originalRequest": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\",\n  \"name\": \"updated Brand\",\n  \"status_id\": true\n//   \"created_at\": \"2023-11-20T12:34:56Z\" // replace with your actual timestamp value\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/brand/655ba0168ed225db7c347301",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"brand",
								"655ba0168ed225db7c347301"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Get All Brand**

```
{
			"name": "Get All Brand",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/brands",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"brands"
					]
				}
			},
			"response": [
				{
					"name": "Get All Brand",
					"originalRequest": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/brands",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"brands"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Detele Brand**

```
{
			"name": "Delete Brand",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "localhost:8080/brand/655ba0168ed225db7c347301",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"brand",
						"655ba0168ed225db7c347301"
					]
				}
			},
			"response": [
				{
					"name": "Delete Brand",
					"originalRequest": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "localhost:8080/brand/655ba0168ed225db7c347301",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"brand",
								"655ba0168ed225db7c347301"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**CreateSupplier**

```
{
			"name": "Create Supplier",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\", // replace with your actual ObjectID value\n  \"email\": \"my@supplier.com\",\n  \"phone\": \"1234567890\",\n  \"status_id\": true,\n  \"is_verified_supplier\": true\n//   \"created_at\": \"2023-11-20T12:34:56Z\" // replace with your actual timestamp value\n}\n",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/supplier",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"supplier"
					]
				}
			},
			"response": [
				{
					"name": "Create Supplier",
					"originalRequest": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\", // replace with your actual ObjectID value\n  \"email\": \"example@supplier.com\",\n  \"phone\": \"1234567890\",\n  \"status_id\": true,\n  \"is_verified_supplier\": true\n//   \"created_at\": \"2023-11-20T12:34:56Z\" // replace with your actual timestamp value\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/supplier",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"supplier"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Get Supplier**

```
{
			"name": "Get Supplier",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/supplier/655c2e288ccec21122733b85",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"supplier",
						"655c2e288ccec21122733b85"
					]
				}
			},
			"response": [
				{
					"name": "Get Supplier",
					"originalRequest": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/supplier/655c2e288ccec21122733b85",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"supplier",
								"655c2e288ccec21122733b85"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Update Supplier**

```
{
			"name": "Update Supplier",
			"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\", // replace with your actual ObjectID value\n  \"email\": \"example@supplier.com\",\n  \"phone\": \"427386528743\",\n  \"status_id\": true,\n  \"is_verified_supplier\": true\n//   \"created_at\": \"2023-11-20T12:34:56Z\" // replace with your actual timestamp value\n}\n",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/supplier/655c2e288ccec21122733b85",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"supplier",
						"655c2e288ccec21122733b85"
					]
				}
			},
			"response": [
				{
					"name": "Update Supplier",
					"originalRequest": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\", // replace with your actual ObjectID value\n  \"email\": \"example@supplier.com\",\n  \"phone\": \"427386528743\",\n  \"status_id\": true,\n  \"is_verified_supplier\": true\n//   \"created_at\": \"2023-11-20T12:34:56Z\" // replace with your actual timestamp value\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/supplier/655c2e288ccec21122733b85",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"supplier",
								"655c2e288ccec21122733b85"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**GetAll Supplier**

```
{
			"name": "Get All Supplier",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/suppliers",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"suppliers"
					]
				}
			},
			"response": [
				{
					"name": "Get All Supplier",
					"originalRequest": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/suppliers",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"suppliers"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Delete Supplier**

```
{
			"name": "Delete Supplier",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "localhost:8080/supplier/655c2e288ccec21122733b85",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"supplier",
						"655c2e288ccec21122733b85"
					]
				}
			},
			"response": [
				{
					"name": "Delete Supplier",
					"originalRequest": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "localhost:8080/supplier/655c2e288ccec21122733b85",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"supplier",
								"655c2e288ccec21122733b85"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Create Product**

```
{
			"name": "Create Product",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\", // replace with your actual ObjectID value\n  \"name\": \"Example Product\",\n  \"description\": \"This is an example product description.\",\n  \"specifications\": \"Example specifications\",\n  \"brand_id\": \"655c41d382916607a487163a\", // replace with your actual Brand ID\n  \"catagory_id\": \"655c41b482916607a4871638\", // replace with your actual Category ID\n  \"supplier_id\": \"655c41fe82916607a487163c\", // replace with your actual Supplier ID\n  \"unit_price\": 100,\n  \"discount_price\": 80,\n  \"tags\": [\"tag1\", \"tag2\"],\n  \"status_id\": true\n}\n",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/product",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"product"
					]
				}
			},
			"response": [
				{
					"name": "Create Product",
					"originalRequest": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n//   \"_id\": \"5f55987bfd86540813b7ec9a\", // replace with your actual ObjectID value\n  \"name\": \"Example Product\",\n  \"description\": \"This is an example product description.\",\n  \"specifications\": \"Example specifications\",\n  \"brand_id\": \"655c41d382916607a487163a\", // replace with your actual Brand ID\n  \"catagory_id\": \"655c41b482916607a4871638\", // replace with your actual Category ID\n  \"supplier_id\": \"655c41fe82916607a487163c\", // replace with your actual Supplier ID\n  \"unit_price\": 100,\n  \"discount_price\": 80,\n  \"tags\": [\"tag1\", \"tag2\"],\n  \"status_id\": true\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/product",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"product"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Get Product**

```
{
			"name": "Get Product",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/product/655c444444452e4523d6af81",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"product",
						"655c444444452e4523d6af81"
					]
				}
			},
			"response": []
		},
```

**Update Product**

```
{
			"name": "Update Product",
			"request": {
				"method": "PUT",
				"header": [],
				"url": {
					"raw": "localhost:8080/category",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"category"
					]
				}
			},
			"response": [
				{
					"name": "Update Product",
					"originalRequest": {
						"method": "PUT",
						"header": [],
						"url": {
							"raw": "localhost:8080/category",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"category"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},

```

**Get All Product**

```
{
			"name": "Get All Product",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/products",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"products"
					]
				}
			},
			"response": [
				{
					"name": "Get All Product",
					"originalRequest": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/products",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"products"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Delete Product**

```
{
			"name": "Delete Product",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "localhost:8080/category",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"category"
					]
				}
			},
			"response": [
				{
					"name": "Delete Product",
					"originalRequest": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "localhost:8080/category",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"category"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Product By Name**

```
{
			"name": "Product By name",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"name\": \"Example Product\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/productbyname",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"productbyname"
					]
				}
			},
			"response": [
				{
					"name": "Product By name",
					"originalRequest": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"name\": \"Example Product\"\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/productbyname",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"productbyname"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Product By Price**

```
{
			"name": "Product By Price",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/productbyprice/100",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"productbyprice",
						"100"
					]
				}
			},
			"response": [
				{
					"name": "Product By Price",
					"originalRequest": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/productbyprice/100",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"productbyprice",
								"100"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Product By Multiple Brand**

```
{
			"name": "Product by Multiple Brand",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"brand_id\":[\"655c41d382916607a487163a\"]\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/productbymulbrand",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"productbymulbrand"
					]
				}
			},
			"response": [
				{
					"name": "Product by Multiple Brand",
					"originalRequest": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"brand_id\":[\"655c41d382916607a487163a\"]\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/productbymulbrand",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"productbymulbrand"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Product By Category**

```
{
			"name": "Product By category",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/productbycategory/655c41b482916607a4871638",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"productbycategory",
						"655c41b482916607a4871638"
					]
				}
			},
			"response": [
				{
					"name": "Product By category",
					"originalRequest": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/productbycategory/655c41b482916607a4871638",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"productbycategory",
								"655c41b482916607a4871638"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Product By Supplier**

```
{
			"name": "Product By supplier",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/productbysupplier/655c41fe82916607a487163c",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"productbysupplier",
						"655c41fe82916607a487163c"
					]
				}
			},
			"response": [
				{
					"name": "Product By supplier",
					"originalRequest": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/productbysupplier/655c41fe82916607a487163c",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"productbysupplier",
								"655c41fe82916607a487163c"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
```

**Product By verified supplier**

```
{
			"name": "Product By verified supplier",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/productbyverifiedsupplier/655c41fe82916607a487163c",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"productbyverifiedsupplier",
						"655c41fe82916607a487163c"
					]
				}
			},
			"response": [
				{
					"name": "Product By verified supplier",
					"originalRequest": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/productbyverifiedsupplier/655c41fe82916607a487163c",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"productbyverifiedsupplier",
								"655c41fe82916607a487163c"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		}

```
