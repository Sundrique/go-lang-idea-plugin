package main
var c = 'a', 'ä', '本', '\t', '\000', '\007', '\377', '\x07', '\xff', '\u12e4', '\U00101234'
/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  VarDeclarationsImpl
    PsiElement(KEYWORD_VAR)('var')
    PsiWhiteSpace(' ')
    VarDeclarationImpl
      LiteralIdentifierImpl
        PsiElement(IDENTIFIER)('c')
      PsiWhiteSpace(' ')
      PsiElement(=)('=')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralCharImpl
          PsiElement(LITERAL_CHAR)(''a'')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralCharImpl
          PsiElement(LITERAL_CHAR)(''ä'')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralCharImpl
          PsiElement(LITERAL_CHAR)(''本'')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralCharImpl
          PsiElement(LITERAL_CHAR)(''\t'')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralCharImpl
          PsiElement(LITERAL_CHAR)(''\000'')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralCharImpl
          PsiElement(LITERAL_CHAR)(''\007'')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralCharImpl
          PsiElement(LITERAL_CHAR)(''\377'')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralCharImpl
          PsiElement(LITERAL_CHAR)(''\x07'')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralCharImpl
          PsiElement(LITERAL_CHAR)(''\xff'')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralCharImpl
          PsiElement(LITERAL_CHAR)(''\u12e4'')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralCharImpl
          PsiElement(LITERAL_CHAR)(''\U00101234'')
