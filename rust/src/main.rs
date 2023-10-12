use lab2::parser::Parser;

fn main() {
    let stdin = std::io::stdin();
    loop {
        let mut buf = String::new();
        let res = stdin.read_line(&mut buf);
        match res {
            Ok(n) => {
                if n == 0 {
                    break;
                }
            }
            Err(err) => panic!("{}", err),
        }
        let mut parser = Parser::default();
        let parsed_result = parser.parse(buf.as_str());
        print!("{}", parsed_result.to_string());
    }
}
