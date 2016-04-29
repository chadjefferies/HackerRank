object Solution {

    def main(args: Array[String]) {
      val input = scala.io.StdIn.readLine().split(" ")
      val n = input(0)
      val k = input(1).toLong

      var sd = n.map(c => c.asDigit*k)
      var res = sd.foldLeft(0L)(_+_).toString
      while (res.length > 1){
        res = superDigit(res).toString()
      }
      println(res)
    }
    
    def superDigit(a:String) : Long = {
      val res = a.map(_.asDigit).toList
      var sum: Long  = 0L
      res.foreach(sum+=_)
      return sum
    }
}
