 object Solution {

    def main(args: Array[String]) {
      val input = readLine().split(" ")
      val n = input(0)
      val k = input(1).toLong

      val superDigit = getSuperDigit(n.map(_.asDigit*k).foldLeft(0L)(_+_).toString)
      println(superDigit)
    }
    
    def getSuperDigit(s:String) : String ={
      if(s.length == 1) return s
      return getSuperDigit(s.map(_.asDigit).foldLeft(0L)(_+_).toString)
    }
}
