package controllers

import javax.inject._
import play.api._
import play.api.mvc._
import play.api.libs.json._
import models.{Product, ProductFormInput}
import scala.collection.mutable.ListBuffer

@Singleton
class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  // Define JSON formatters for the Product class
  implicit val productWrites: OWrites[Product] = Json.writes[Product]
  implicit val productReads: Reads[Product] = Json.reads[Product]

  // Form input model without ID
  implicit val productFormInputReads: Reads[ProductFormInput] = Json.reads[ProductFormInput]

  // In-memory product database
  private val products = ListBuffer(
    Product(1, "Laptop", "High-performance laptop", 1299.99),
    Product(2, "Smartphone", "Latest smartphone model", 899.99),
    Product(3, "Headphones", "Noise-cancelling headphones", 249.99),
    Product(4, "Tablet", "Professional tablet with stylus", 799.99),
    Product(5, "Smartwatch", "Fitness tracking smartwatch", 199.99)
  )


  // CREATE - Add a new product (using ProductFormInput without ID)
  def create = Action(parse.json) { (request: Request[JsValue]) =>
    request.body.validate[ProductFormInput].fold(
      errors => {
        BadRequest(Json.obj("status" -> "error", "message" -> JsError.toJson(errors)))
      },
      productInput => {
        // Generate a new ID
        val newId = if (products.isEmpty) 1 else products.map(_.id).max + 1

        // Create a new product with the generated ID
        val newProduct = Product(
          id = newId,
          name = productInput.name,
          description = productInput.description,
          price = productInput.price
        )

        products += newProduct
        Created(Json.toJson(newProduct))
      }
    )
  }

  // READ - Get all products
  def getAll = Action { (request: Request[AnyContent]) =>
    Ok(Json.toJson(products))
  }

  // READ - Get a specific product by ID
  def getById(id: Long) = Action { (request: Request[AnyContent]) =>
    products.find(_.id == id) match {
      case Some(product) => Ok(Json.toJson(product))
      case None => NotFound(Json.obj("status" -> "error", "message" -> s"Product with ID $id not found"))
    }
  }

  // UPDATE - Update an existing product
  def update(id: Long) = Action(parse.json) { (request: Request[JsValue]) =>
    request.body.validate[ProductFormInput].fold(
      errors => {
        BadRequest(Json.obj("status" -> "error", "message" -> JsError.toJson(errors)))
      },
      productInput => {
        products.indexWhere(_.id == id) match {
          case -1 => NotFound(Json.obj("status" -> "error", "message" -> s"Product with ID $id not found"))
          case index =>
            // Create updated product with the path ID parameter and form input data
            val updatedProduct = Product(
              id = id,
              name = productInput.name,
              description = productInput.description,
              price = productInput.price
            )

            products(index) = updatedProduct
            Ok(Json.toJson(updatedProduct))
        }
      }
    )
  }

  // DELETE - Remove a product
  def delete(id: Long) = Action { (request: Request[AnyContent]) =>
    products.indexWhere(_.id == id) match {
      case -1 => NotFound(Json.obj("status" -> "error", "message" -> s"Product with ID $id not found"))
      case index =>
        val deleted = products.remove(index)
        Ok(Json.obj("status" -> "success", "message" -> s"Product '${deleted.name}' deleted"))
    }
  }

  // Original index method
  def index() = Action { (request: Request[AnyContent]) =>
    Ok(views.html.index())
  }
}
