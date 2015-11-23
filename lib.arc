(assign mac (annotate 'macro (fn (name args . forms) `(assign ,name (annotate 'macro (fn ,args ,@forms))))))

(mac def (name args . forms)
  `(assign ,name (fn ,args ,@forms)))

(def prn (arg . args)
  (disp arg)
  (if args
      (prn . args)
    ((fn ()
      (disp #\newline)
      arg))))

(mac let (var val . forms)
  `((fn (,var) ,@forms) ,val))

(mac = (place value)
  (if (is (type place) 'cons)
      `(sref ,@place ,value)
    `(assign ,place , value)))

(def no (x)
  (if x nil t))

(def list args
  args)

(mac do forms
  `((fn () ,@ forms)))

(mac whilet (var test . body)
  (w/uniq (gf gp)
    ``((rfn ,gf (,gp)
        (let ,var, gp
          (when ,var ,@body (,gf ,test))))
       ,test)))

(= current-stdout (stdout))
(= current-stdin (stdin))
